package scraping

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jcasado94/nats-points/mongo/drivers"
	"github.com/jcasado94/nats-points/mongo/entity"
)

const (
	name      = "theguardian"
	sharesUrl = "https://api.nextgen.guardianapps.co.uk/sharecount/%s.json"
)

type TheGuardianScraper struct {
	client http.Client
	md     *drivers.MongoDriver
}

func NewTheGuardianScraper() (TheGuardianScraper, error) {
	md, err := drivers.NewMongoDriver()
	if err != nil {
		return TheGuardianScraper{}, err
	}
	return TheGuardianScraper{
		client: http.Client{},
		md:     &md,
	}, nil
}

func (s *TheGuardianScraper) GetAllNews(countryName string) ([]entity.Article, error) {
	articles := make([]entity.Article, 0)

	urlPattern, err := s.md.GetArticlesUrl(countryName, name)
	if err != nil {
		return nil, err
	}
	i := 1
	finished := false
	for {
		url := fmt.Sprintf(urlPattern, i)
		resp, err := s.client.Get(url)
		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return nil, err
		}
		// concurrent
		doc.Find(".fc-container--tag").Each(func(i int, sel *goquery.Selection) {
			dateString, exists := sel.Find(".fc-container__header time").Attr("datetime")
			if !exists {
				log.Print("couldn't find date attribute")
				return
			}
			dateSlice := strings.Split(dateString, "-")
			year, err := strconv.Atoi(dateSlice[0])
			month, err := strconv.Atoi(dateSlice[1])
			if err != nil {
				log.Printf("couldn't parse date %s", dateString)
				return
			}
			date := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
			if time.Now().Sub(date) <= maxNewsAge {
				// concurrent
				sel.Find(".fc-item__link").Each(func(j int, articleSel *goquery.Selection) {
					var article entity.Article
					articleLink, _ := articleSel.Attr("href")
					article.Url = articleLink
					resp, err = s.client.Get(articleLink)
					if err != nil {
						log.Printf("couldn't load url %s", articleLink)
						return
					}
					defer resp.Body.Close()
					articleDoc, err := goquery.NewDocumentFromReader(resp.Body)
					if err != nil {
						log.Printf("couldn't parse url %s", articleLink)
						return
					}

					if keywordsSelection, exists := findAttrElem(articleDoc, "meta", "name", "keywords"); exists {
						if keywords, exist := keywordsSelection.Attr("content"); exist {
							keywordsSlice := strings.Split(keywords, ",")
							article.Tags = keywordsSlice
						} else {
							log.Printf("couldn't find keywords for %s", articleLink)
						}
					} else {
						log.Printf("couldn't find keywords for %s", articleLink)
					}

					sharesJson, err := getSharesJson(articleLink, &s.client)
					if err != nil {
						log.Print(err.Error())
					} else {
						article.Shares = sharesJson.ShareCount
					}

					articles = append(articles, article)
				})
			} else {
				finished = true
			}
		})
		i++
		if finished {
			break
		}
	}

	return articles, nil

}

func findAttrElem(doc *goquery.Document, elem, attrKey, attrValue string) (node *goquery.Selection, found bool) {
	var res *goquery.Selection
	doc.Find(elem).Each(func(i int, s *goquery.Selection) {
		if attr, exists := s.Attr(attrKey); exists {
			if attr == attrValue {
				res = s
			}
		}
	})
	if res == nil {
		return res, false
	}
	return res, true
}

type SharesJson struct {
	ShareCount int `json:"share_count"`
}

func getSharesJson(url string, client *http.Client) (SharesJson, error) {
	resp, err := client.Get(fmt.Sprintf(sharesUrl, getPath(url)))
	if err != nil {
		return SharesJson{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return SharesJson{}, err
	}
	var sharesJson SharesJson
	err = json.Unmarshal(body, &sharesJson)
	if err != nil {
		return SharesJson{}, err
	}
	return sharesJson, nil
}

func getPath(url string) string {
	split := strings.Split(url, "theguardian.com/")
	return split[len(split)-1]
}

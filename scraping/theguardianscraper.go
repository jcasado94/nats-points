package scraping

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/jcasado94/nats-points/mongo/drivers"
	"github.com/jcasado94/nats-points/mongo/entity"
)

const name = "theguardian"

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

func (s *TheGuardianScraper) GetAllNews(countryName string) ([]entity.News, error) {
	news := make([]entity.Articles, 0)

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
				err = errors.New("couldn't find date attribute")
				return
			}
			dateSlice := strings.Split(dateString, "-")
			year, err := strconv.Atoi(dateSlice[0])
			month, err := strconv.Atoi(dateSlice[1])
			if err != nil {
				err = errors.New("couldn't parse date")
				return
			}
			date := time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
			if time.Now().Sub(date) <= maxNewsAge {
				// concurrent
				sel.Find(".fc-item__link").Each(func(j int, articleSel *goquery.Selection) {
					articleLink, _ := sel.Attr("href")
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
					if keywords, exist := articleDoc.Find("meta[name='keywords']"); exist {
						keywordsSlice, err := strings.Split(keywords, ",")
						
					} else {
						log.Printf("couldn't find keywords for %s", articleLink)
					}
				})
			} else {
				finished = true
				return
			}
		})
		if finished {
			break
		} else {
			return nil, err
		}
		i++
	}

	return news, nil

}

package invalidator

import (
	"github.com/jcasado94/nats-points/mongo/drivers"
	"github.com/jcasado94/nats-points/mongo/entity"
	"github.com/jcasado94/nats-points/mongo/model"
	"github.com/jcasado94/nats-points/scraping"
)

type Invalidator struct {
	scrapers []scraping.Scraper
	driver   *drivers.MongoDriver
}

func NewInvalidator(md *drivers.MongoDriver) (Invalidator, error) {
	tg, err := scraping.NewTheGuardianScraper()
	if err != nil {
		return Invalidator{}, err
	}
	return Invalidator{
		scrapers: []scraping.Scraper{
			tg,
		},
		driver: md,
	}, nil
}

func (inv *Invalidator) InvalidateAllCountryArticles(countryName string) error {
	allArticles := make([]entity.Article, 0)
	for _, s := range inv.scrapers {
		articles, err := s.GetAllArticles(countryName)
		if err != nil {
			return err
		}
		allArticles = append(allArticles, articles...)
	}
	oldIds, err := inv.driver.DeleteAllCountryArticles(countryName)
	if err != nil {
		return err
	}
	err = inv.driver.DeleteArticles(oldIds)
	if err != nil {
		return err
	}
	allModelArticles := make([]model.ArticleModel, 0)
	for _, article := range allArticles {
		allModelArticles = append(allModelArticles, model.NewArticleModel(&article))
	}
	inv.driver.AddArticles(allModelArticles)
	insertedModelArticles := make([]model.ArticleModel, 0)
	for _, art := range allModelArticles {
		insertedArticle, err := inv.driver.GetArticleByUrl(art.Url)
		if err != nil {
			return err
		}
		insertedModelArticles = append(insertedModelArticles, insertedArticle)
	}
	return inv.driver.AddAllArticles(countryName, insertedModelArticles)
}

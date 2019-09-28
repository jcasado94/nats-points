package invalidator

import (
	"github.com/jcasado94/nats-points/mongo/drivers"
	"github.com/jcasado94/nats-points/mongo/entity"
	"github.com/jcasado94/nats-points/scraping"
)

type Invalidator struct {
	scrapers []scraping.Scraper
	driver   *drivers.MongoDriver
}

func NewInvalidator() (Invalidator, error) {
	tg, err := scraping.NewTheGuardianScraper()
	if err != nil {
		return Invalidator{}, err
	}
	md, err := drivers.NewMongoDriver()
	if err != nil {
		return Invalidator{}, err
	}
	return Invalidator{
		scrapers: []scraping.Scraper{
			tg,
		},
		driver: &md,
	}, nil
}

func (inv *Invalidator) InvalidateAllArticles(countryName string) error {
	allArticles := make([]entity.Article, 0)
	for _, s := range inv.scrapers {
		articles, err := s.GetAllArticles(countryName)
		if err != nil {
			return err
		}
		allArticles = append(allArticles, articles...)
	}
	err := inv.driver.DeleteAllArticles()
	if err != nil {
		return err
	}
	err = inv.driver.InsertAllArticles(allArticles)
	if err != nil {
		return err
	}
	err = inv.driver.DeleteAllCountryArticles(countryName)
	if err != nil {
		return err
	}
	modelArticles, err := inv.driver.GetAllArticles()
	if err != nil {
		return err
	}
	return inv.driver.AddAllArticles(countryName, modelArticles)
}

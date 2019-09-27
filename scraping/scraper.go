package scraping

import (
	"time"

	"github.com/jcasado94/nats-points/mongo/entity"
)

var maxNewsAge, _ = time.ParseDuration("2190h")

type Scraper interface {
	GetAllNews(countryName string) []entity.Articles
}

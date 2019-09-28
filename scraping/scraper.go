package scraping

import (
	"time"

	"github.com/jcasado94/nats-points/mongo/entity"
)

var maxNewsAge, _ = time.ParseDuration("730h")

type Scraper interface {
	GetAllArticles(countryName string) ([]entity.Article, error)
}

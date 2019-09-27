package drivers

import (
	"github.com/jcasado94/nats-points/mongo"
	mongoService "github.com/jcasado94/nats-points/mongo/service"
)

const (
	mongoEndpoint            = "localhost:27017"
	mongoDb                  = "localup"
	mongoCountriesCollection = "countries"
	mongoArticlesCollection  = "articles"
)

type MongoDriver struct {
	session         *mongo.Session
	countryService  *mongoService.CountryService
	articlesService *mongoService.ArticleService
}

func NewMongoDriver() (MongoDriver, error) {
	session, err := mongo.NewSession(mongoEndpoint)
	if err != nil {
		return MongoDriver{}, err
	}
	return MongoDriver{
		session:         session,
		countryService:  mongoService.NewCountryService(session, mongoDb, mongoCountriesCollection),
		articlesService: mongoService.NewArticleService(session, mongoDb, mongoArticlesCollection),
	}, nil
}

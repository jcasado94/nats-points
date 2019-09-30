package service

import (
	"errors"

	"github.com/jcasado94/nats-points/mongo"
	"github.com/jcasado94/nats-points/mongo/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CountryService struct {
	collection *mgo.Collection
}

func NewCountryService(session *mongo.Session, dbName, colName string) *CountryService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(entity.CountryModelIndex())
	return &CountryService{collection}
}

func (cs *CountryService) GetCountry(countryName string) (entity.Country, error) {
	query := map[string]string{"name": countryName}
	var c entity.Country
	err := cs.collection.Find(query).One(&c)
	if err != nil {
		return entity.Country{}, err
	}
	return c, nil
}

func (cs *CountryService) GetArticlesUrl(countryName string, provider string) (string, error) {
	country, err := cs.GetCountry(countryName)
	if err != nil {
		return "", err
	}
	if _, exists := country.ArticlesUrls[provider]; !exists {
		return "", errors.New("no newspaper found")
	}
	return country.ArticlesUrls[provider], nil
}

func (cs *CountryService) GetAllArticlesIDs(countryName string) ([]bson.ObjectId, error) {
	country, err := cs.GetCountry(countryName)
	if err != nil {
		return nil, err
	}
	bsonIDs := country.GetAllArticleIDs()
	return bsonIDs, nil
}

func objectIdSliceToString(sl []bson.ObjectId) []string {
	res := make([]string, 0)
	for _, obj := range sl {
		res = append(res, obj.String())
	}
	return res
}

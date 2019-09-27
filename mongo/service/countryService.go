package service

import (
	"errors"

	"github.com/jcasado94/nats-points/mongo"
	"github.com/jcasado94/nats-points/mongo/entity"
	"github.com/jcasado94/nats-points/mongo/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CountryService struct {
	collection *mgo.Collection
}

func NewCountryService(session *mongo.Session, dbName, colName string) *CountryService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(model.CountryModelIndex())
	return &CountryService{collection}
}

func (cs *CountryService) GetCountry(countryName string) (entity.Country, error) {
	query := map[string]string{"name": countryName}
	var cm model.CountryModel
	err := cs.collection.Find(query).One(&cm)
	if err != nil {
		return entity.Country{}, err
	}
	return entity.NewCountry(cm.Info.Population, objectIdSliceToString(cm.Articles.Environment), objectIdSliceToString(cm.Articles.Politics), cm.ArticlesUrls), nil
}

func (cs *CountryService) GetArticlesUrl(countryName, newspaper string) (string, error) {
	country, err := cs.GetCountry(countryName)
	if err != nil {
		return "", err
	}
	if _, exists := country.ArticlesUrls[newspaper]; !exists {
		return "", errors.New("no newspaper found")
	}
	return country.ArticlesUrls[newspaper], nil
}

func objectIdSliceToString(sl []bson.ObjectId) []string {
	res := make([]string, 0)
	for _, obj := range sl {
		res = append(res, obj.String())
	}
	return res
}

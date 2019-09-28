package service

import (
	"errors"

	"github.com/jcasado94/nats-points/mongo"
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

func (cs *CountryService) GetCountry(countryName string) (model.CountryModel, error) {
	query := map[string]string{"name": countryName}
	var cm model.CountryModel
	err := cs.collection.Find(query).One(&cm)
	if err != nil {
		return model.CountryModel{}, err
	}
	return cm, nil
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

func (cs *CountryService) DeleteAllArticles(countryName string) error {
	country, err := cs.GetCountry(countryName)
	if err != nil {
		return err
	}
	country.Articles = model.Articles{
		Environment: []bson.ObjectId{},
		Politics:    []bson.ObjectId{},
	}
	return cs.collection.Update(map[string]string{"name": countryName}, country)
}

func (cs *CountryService) AddAllArticles(countryName string, modelArticles []model.ArticleModel) error {
	country, err := cs.GetCountry(countryName)
	if err != nil {
		return err
	}
	country.PopulateArticles(modelArticles)
	return cs.collection.Update(map[string]string{"name": countryName}, country)
}

func objectIdSliceToString(sl []bson.ObjectId) []string {
	res := make([]string, 0)
	for _, obj := range sl {
		res = append(res, obj.String())
	}
	return res
}

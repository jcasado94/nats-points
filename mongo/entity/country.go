package entity

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var articleTags = []string{"environment", "politics", "society", "sport", "business", "culture"}

type Country struct {
	ID           bson.ObjectId              `json:"id"`
	Name         string                     `json:"name"`
	Info         Information                `json:"information"`
	Articles     map[string][]bson.ObjectId `json:"articles"`
	ArticlesUrls map[string]string          `json:"articlesUrls"`
}

func (c *Country) GetAllArticleIDs() []bson.ObjectId {
	ids := make([]bson.ObjectId, 0)
	for _, tagArticles := range c.Articles {
		ids = append(ids, tagArticles...)
	}
	return ids
}

type Information struct {
	Population int                `json:"population"`
	Area       int                `json:"area"`
	Capital    string             `json:"capital"`
	Currency   string             `json:"currency"`
	Conversion map[string]float64 `json:"conversion"`
}

func CountryModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"ID"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

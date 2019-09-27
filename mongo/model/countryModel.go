package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CountryModel struct {
	ID           bson.ObjectId     `bson:"_id,omitempty"`
	Info         Information       `bson:"information"`
	Articles     Articles          `bson:"articles"`
	ArticlesUrls map[string]string `bson:"articlesUrls"`
}

type Information struct {
	Population int `bson:"population"`
}

type Articles struct {
	Environment []bson.ObjectId `bson:"environment"`
	Politics    []bson.ObjectId `bson:"politics"`
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

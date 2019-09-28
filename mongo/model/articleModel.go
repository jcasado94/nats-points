package model

import (
	"github.com/jcasado94/nats-points/mongo/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ArticleModel struct {
	ID     bson.ObjectId `bson:"_id,omitempty"`
	Url    string        `bson:"url"`
	Title  string        `bson:"title"`
	Img    string        `bson:"img"`
	Tags   []string      `bson:"tags"`
	Shares int           `bson:"shares"`
}

func NewArticleModel(a *entity.Article) ArticleModel {
	return ArticleModel{
		Url:    a.Url,
		Title:  a.Title,
		Img:    a.Img,
		Tags:   a.Tags,
		Shares: a.Shares,
	}
}

func ArticleModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"ID", "Url"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

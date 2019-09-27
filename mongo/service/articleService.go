package service

import (
	"github.com/jcasado94/nats-points/mongo"
	"github.com/jcasado94/nats-points/mongo/entity"
	"github.com/jcasado94/nats-points/mongo/model"
	"gopkg.in/mgo.v2"
)

type ArticleService struct {
	collection *mgo.Collection
}

func NewArticleService(session *mongo.Session, dbName, colName string) *ArticleService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(model.ArticleModelIndex())
	return &ArticleService{collection}
}

func (as *ArticleService) InsertArticle(a *entity.Article) error {
	return as.collection.Insert(a)
}

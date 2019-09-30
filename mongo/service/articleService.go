package service

import (
	"github.com/jcasado94/nats-points/mongo"
	"github.com/jcasado94/nats-points/mongo/entity"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ArticleService struct {
	collection *mgo.Collection
}

func NewArticleService(session *mongo.Session, dbName, colName string) *ArticleService {
	collection := session.GetCollection(dbName, colName)
	collection.EnsureIndex(entity.ArticleModelIndex())
	return &ArticleService{collection}
}

func (as *ArticleService) InsertArticle(a *entity.Article) error {
	return as.collection.Insert(a)
}

func (as *ArticleService) DeleteAllArticles() error {
	_, err := as.collection.RemoveAll(nil)
	return err
}

func (as *ArticleService) DeleteArticles(ids []bson.ObjectId) error {
	for _, id := range ids {
		err := as.collection.Remove(map[string]bson.ObjectId{"_id": id})
		if err != nil {
			return err
		}
	}
	return nil
}

func (as *ArticleService) GetArticles(ids []bson.ObjectId) []entity.Article {
	articles := make([]entity.Article, 0)
	as.collection.Find(bson.M{
		"_id": bson.M{
			"$in": ids,
		},
	}).All(&articles)
	return articles
}

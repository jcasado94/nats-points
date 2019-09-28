package service

import (
	"github.com/jcasado94/nats-points/mongo"
	"github.com/jcasado94/nats-points/mongo/entity"
	"github.com/jcasado94/nats-points/mongo/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

func (as *ArticleService) DeleteAllArticles() error {
	_, err := as.collection.RemoveAll(nil)
	return err
}

func (as *ArticleService) InsertAllArticles(articles []model.ArticleModel) error {
	var err error
	for _, art := range articles {
		err = as.collection.Insert(art)
	}
	return err
}

func (as *ArticleService) GetAllArticles() ([]model.ArticleModel, error) {
	var articles []model.ArticleModel
	err := as.collection.Find(nil).All(&articles)
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (as *ArticleService) GetAllArticlesMapped() (map[bson.ObjectId]model.ArticleModel, error) {
	modelArticles, err := as.GetAllArticles()
	if err != nil {
		return nil, err
	}
	res := make(map[bson.ObjectId]model.ArticleModel)
	for _, art := range modelArticles {
		res[art.ID] = art
	}
	return res, nil
}

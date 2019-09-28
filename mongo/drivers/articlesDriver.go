package drivers

import (
	"log"

	"github.com/jcasado94/nats-points/mongo/entity"
	"github.com/jcasado94/nats-points/mongo/model"
	"gopkg.in/mgo.v2/bson"
)

func (md *MongoDriver) InsertArticle(a *entity.Article) error {
	return md.articlesService.InsertArticle(a)
}

func (md *MongoDriver) DeleteAllArticles() error {
	return md.articlesService.DeleteAllArticles()
}

func (md *MongoDriver) AddArticles(modelArticles []model.ArticleModel) {
	md.articlesService.InsertAllArticles(modelArticles)
}

func (md *MongoDriver) GetArticleByUrl(url string) (model.ArticleModel, error) {
	return md.articlesService.GetArticleByUrl(url)
}

func (md *MongoDriver) InsertAllArticles(as []entity.Article) error {
	err := md.articlesService.DeleteAllArticles()
	if err != nil {
		log.Print(err.Error())
	}
	articlesModel := make([]model.ArticleModel, 0)
	for _, a := range as {
		articlesModel = append(articlesModel, model.NewArticleModel(&a))
	}
	return md.articlesService.InsertAllArticles(articlesModel)
}

func (md *MongoDriver) GetAllArticles() ([]model.ArticleModel, error) {
	return md.articlesService.GetAllArticles()
}

func (md *MongoDriver) GetAllArticlesMapped() (map[bson.ObjectId]model.ArticleModel, error) {
	return md.articlesService.GetAllArticlesMapped()
}

func (md *MongoDriver) DeleteArticles(ids []bson.ObjectId) error {
	return md.articlesService.DeleteArticles(ids)
}

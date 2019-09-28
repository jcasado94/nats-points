package drivers

import (
	"log"

	"github.com/jcasado94/nats-points/mongo/entity"
	"github.com/jcasado94/nats-points/mongo/model"
)

func (md *MongoDriver) InsertArticle(a *entity.Article) error {
	return md.articlesService.InsertArticle(a)
}

func (md *MongoDriver) DeleteAllArticles() error {
	return md.articlesService.DeleteAllArticles()
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

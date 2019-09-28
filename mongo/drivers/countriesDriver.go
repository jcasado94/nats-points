package drivers

import (
	"github.com/jcasado94/nats-points/mongo/entity"
	"github.com/jcasado94/nats-points/mongo/model"
)

func (md *MongoDriver) GetArticlesUrl(countryName, newspaper string) (string, error) {
	return md.countryService.GetArticlesUrl(countryName, newspaper)
}

func (md *MongoDriver) DeleteAllCountryArticles(countryName string) error {
	return md.countryService.DeleteAllArticles(countryName)
}

func (md *MongoDriver) AddAllArticles(countryName string, modelArticles []model.ArticleModel) error {
	return md.countryService.AddAllArticles(countryName, modelArticles)
}

func (md *MongoDriver) GetAllCountryArticles(countryName string) (model.Articles, error) {
	return md.countryService.GetAllArticles(countryName)
}

func (md *MongoDriver) GetAllCountryResultArticles(countryName string) (entity.ResultArticles, error) {
	articles, err := md.GetAllCountryArticles(countryName)
	if err != nil {
		return entity.ResultArticles{}, err
	}
	mappedArticles, err := md.articlesService.GetAllArticlesMapped()
	if err != nil {
		return entity.ResultArticles{}, err
	}
	res := entity.NewResultArticles()
	for _, artId := range articles.Environment {
		modelArt := mappedArticles[artId]
		art := model.ArticleFromModel(&modelArt)
		res.Environment.Add(&art)
	}
	res.Environment = res.Environment.GetElements()
	for _, artId := range articles.Politics {
		modelArt := mappedArticles[artId]
		art := model.ArticleFromModel(&modelArt)
		res.Politics.Add(&art)
	}
	res.Politics = res.Politics.GetElements()
	for _, artId := range articles.Society {
		modelArt := mappedArticles[artId]
		art := model.ArticleFromModel(&modelArt)
		res.Society.Add(&art)
	}
	res.Society = res.Society.GetElements()
	for _, artId := range articles.Sports {
		modelArt := mappedArticles[artId]
		art := model.ArticleFromModel(&modelArt)
		res.Sports.Add(&art)
	}
	res.Sports = res.Sports.GetElements()
	for _, artId := range articles.Business {
		modelArt := mappedArticles[artId]
		art := model.ArticleFromModel(&modelArt)
		res.Business.Add(&art)
	}
	res.Business = res.Business.GetElements()
	for _, artId := range articles.Culture {
		modelArt := mappedArticles[artId]
		art := model.ArticleFromModel(&modelArt)
		res.Culture.Add(&art)
	}
	res.Culture = res.Culture.GetElements()

	return res, nil

}

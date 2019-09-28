package drivers

import "github.com/jcasado94/nats-points/mongo/model"

func (md *MongoDriver) GetArticlesUrl(countryName, newspaper string) (string, error) {
	return md.countryService.GetArticlesUrl(countryName, newspaper)
}

func (md *MongoDriver) DeleteAllCountryArticles(countryName string) error {
	return md.countryService.DeleteAllArticles(countryName)
}

func (md *MongoDriver) AddAllArticles(countryName string, modelArticles []model.ArticleModel) error {
	return md.countryService.AddAllArticles(countryName, modelArticles)
}

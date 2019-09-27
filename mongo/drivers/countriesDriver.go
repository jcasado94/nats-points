package drivers

import "github.com/jcasado94/nats-points/mongo/entity"

func (md *MongoDriver) GetCountry(countryName string) (entity.Country, error) {
	return md.countryService.GetCountry(countryName)
}

func (md *MongoDriver) GetArticlesUrl(countryName, newspaper string) (string, error) {
	return md.countryService.GetArticlesUrl(countryName, newspaper)
}

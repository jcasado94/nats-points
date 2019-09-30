package drivers

import (
	"github.com/jcasado94/nats-points/mongo/entity"
	"gopkg.in/mgo.v2/bson"
)

func (md *MongoDriver) GetArticles(ids []bson.ObjectId) []entity.Article {
	return md.articlesService.GetArticles(ids)
}

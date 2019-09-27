package drivers

import "github.com/jcasado94/nats-points/mongo/entity"

func (md *MongoDriver) InsertArticle(a *entity.Article) error {
	return md.articlesService.InsertArticle(a)
}

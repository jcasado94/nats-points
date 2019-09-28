package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CountryModel struct {
	ID           bson.ObjectId     `bson:"_id,omitempty"`
	Name         string            `bson:"name"`
	Info         Information       `bson:"information"`
	Articles     Articles          `bson:"articles"`
	ArticlesUrls map[string]string `bson:"articlesUrls"`
}

func (cm *CountryModel) PopulateArticles(arts []ArticleModel) {
	for _, art := range arts {
		for _, tag := range art.Tags {
			tag = tagMapping[tag]
			added := true
			switch tag {
			case "politics":
				cm.Articles.Politics = append(cm.Articles.Politics, art.ID)
			case "environment":
				cm.Articles.Environment = append(cm.Articles.Environment, art.ID)
			case "culture":
				cm.Articles.Culture = append(cm.Articles.Culture, art.ID)
			case "sports":
				cm.Articles.Sports = append(cm.Articles.Sports, art.ID)
			case "business":
				cm.Articles.Business = append(cm.Articles.Business, art.ID)
			case "society":
				cm.Articles.Society = append(cm.Articles.Society, art.ID)
			default:
				added = false
			}
			if added {
				break
			}
		}
	}
}

type Information struct {
	Population int `bson:"population"`
}

type Articles struct {
	Environment []bson.ObjectId `bson:"environment"`
	Politics    []bson.ObjectId `bson:"politics"`
	Society     []bson.ObjectId `bson:"society"`
	Sports      []bson.ObjectId `bson:"sports"`
	Business    []bson.ObjectId `bson:"business"`
	Culture     []bson.ObjectId `bson:"culture"`
}

func CountryModelIndex() mgo.Index {
	return mgo.Index{
		Key:        []string{"ID"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
}

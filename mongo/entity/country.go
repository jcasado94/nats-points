package entity

type Country struct {
	Name         string
	Info         Information       `json:"information"`
	Articles     Articles          `json:"articles"`
	ArticlesUrls map[string]string `json:"articlesUrls"`
}

func NewCountry(population int, environmentNews []string, politicsNews []string, articlesUrls map[string]string) Country {
	return Country{
		Info:         newInformation(population),
		Articles:     newNews(environmentNews, politicsNews),
		ArticlesUrls: articlesUrls,
	}
}

type CountryService interface {
	GetCountry(countryName string) Country
	GetArticlesUrl(countryName, newspaper string) string
}

type Articles struct {
	Environment []string `json:"environment"`
	Politics    []string `json:"politics"`
}

func newNews(environmentNews, politicsNews []string) Articles {
	return Articles{
		Environment: environmentNews,
		Politics:    politicsNews,
	}
}

type Information struct {
	Population int `json:"population"`
}

func newInformation(population int) Information {
	return Information{
		Population: population,
	}
}

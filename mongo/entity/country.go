package entity

type Country struct {
	Name         string
	Info         Information       `json:"information"`
	Articles     Articles          `json:"articles"`
	ArticlesUrls map[string]string `json:"articlesUrls"`
}

func NewCountry(population int, environmentNews, politicsNews, societyNews, sportsNews, businessNews, cultureNews []string, articlesUrls map[string]string) Country {
	return Country{
		Info:         newInformation(population),
		Articles:     newArticles(environmentNews, politicsNews, societyNews, sportsNews, businessNews, cultureNews),
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
	Society     []string `json:"society"`
	Sports      []string `json:"sports"`
	Business    []string `json:"business"`
	Culture     []string `json:"culture"`
}

func newArticles(environmentNews, politicsNews, societyNews, sportsNews, businessNews, cultureNews []string) Articles {
	return Articles{
		Environment: environmentNews,
		Politics:    politicsNews,
		Society:     societyNews,
		Sports:      sportsNews,
		Business:    businessNews,
		Culture:     cultureNews,
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

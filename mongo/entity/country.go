package entity

import "container/heap"

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

type ResultArticlesTagged struct {
	Environment, Politics, Society, Sports, Business, Culture SortableArticles
}

func NewResultArticles() ResultArticlesTagged {
	res := ResultArticlesTagged{
		Environment: make(SortableArticles, 0),
		Politics:    make(SortableArticles, 0),
		Society:     make(SortableArticles, 0),
		Sports:      make(SortableArticles, 0),
		Business:    make(SortableArticles, 0),
		Culture:     make(SortableArticles, 0),
	}
	heap.Init(&res.Environment)
	heap.Init(&res.Politics)
	heap.Init(&res.Society)
	heap.Init(&res.Sports)
	heap.Init(&res.Business)
	heap.Init(&res.Culture)

	return res
}

func (ra *ResultArticlesTagged) MergeTags() SortableArticles {
	res := make(SortableArticles, 0)
	heap.Init(&res)
	for _, art := range ra.Environment {
		heap.Push(&res, art)
	}
	for _, art := range ra.Politics {
		heap.Push(&res, art)
	}
	for _, art := range ra.Society {
		heap.Push(&res, art)
	}
	for _, art := range ra.Sports {
		heap.Push(&res, art)
	}
	for _, art := range ra.Business {
		heap.Push(&res, art)
	}
	for _, art := range ra.Culture {
		heap.Push(&res, art)
	}

	return res
}

type SortableArticles []Article

func (h SortableArticles) GetElements() []Article {
	res := make([]Article, 0)
	for len(h) != 0 {
		res = append(res, heap.Pop(&h).(Article))
	}
	return res
}

func (h *SortableArticles) Add(art *Article) {
	heap.Push(h, *art)
}

func (h SortableArticles) Len() int           { return len(h) }
func (h SortableArticles) Less(i, j int) bool { return h[i].Shares > h[j].Shares }
func (h SortableArticles) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *SortableArticles) Push(x interface{}) {
	*h = append(*h, x.(Article))
}

func (h *SortableArticles) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Information struct {
	Population int `json:"population"`
}

func newInformation(population int) Information {
	return Information{
		Population: population,
	}
}

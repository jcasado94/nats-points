package entity

type Article struct {
	Url    string   `json:"url"`
	Tags   []string `json:"tags"`
	Shares int      `json:"shares"`
}

func NewArticle(url string, tags []string, shares int) Article {
	return Article{
		Url:    url,
		Tags:   tags,
		Shares: shares,
	}
}

type ArticleService interface {
	InsertArticle(a *Article) error
}

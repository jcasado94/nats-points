package entity

type Article struct {
	Url    string   `json:"url"`
	Title  string   `json:"title"`
	Img    string   `json:"img"`
	Tags   []string `json:"tags"`
	Shares int      `json:"shares"`
}

func NewArticle(url, title, img string, tags []string, shares int) Article {
	return Article{
		Url:    url,
		Title:  title,
		Img:    img,
		Tags:   tags,
		Shares: shares,
	}
}

type ArticleService interface {
	InsertArticle(a *Article) error
	DeleteAllArticles() error
	InsertAllArticles(articles []Article) error
}

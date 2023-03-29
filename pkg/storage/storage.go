package storage

type Article struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	AuthorName  string
	CreatedAt   int64
	PublishedAt int64
}

type DBInterface interface {
	ArticlesList() ([]Article, error)
	AddArticle(Article) error
	UpdateArticle(Article) error
	DeleteArticle(Article) error
}

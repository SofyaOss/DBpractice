package memdb

import "skillfactory/DBpractice/pkg/storage"

type memDB struct {
	data []storage.Article
}

func New() *memDB {
	return new(memDB)
}

func (s *memDB) Articles() ([]storage.Article, error) {
	return articles, nil
}

func (s *memDB) AddArticle(storage.Article) error {
	return nil
}
func (s *memDB) UpdateArticle(storage.Article) error {
	return nil
}
func (s *memDB) DeleteArticle(storage.Article) error {
	return nil
}

var articles = []storage.Article{
	{
		ID:      1,
		Title:   "title 1",
		Content: "content 1",
	},
	{
		ID:      2,
		Title:   "title 2",
		Content: "content 2",
	},
}

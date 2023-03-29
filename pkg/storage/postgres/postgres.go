package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"skillfactory/DBpractice/pkg/storage"
)

const (
	databaseUrl string = "user=postgres password=Keks17sql dbname=articlesDB sslmode=disable" // "postgres://postgres:postgres@127.0.0.1:8081/articlesDB"
)

type postgresDB struct {
	ctx  context.Context
	pool *pgxpool.Pool
}

func New() (*postgresDB, error) {
	pool, err := pgxpool.Connect(context.Background(), databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return &postgresDB{ctx: context.Background(), pool: pool}, nil
}

func (p *postgresDB) ArticlesList() ([]storage.Article, error) { // получение всех публикаций
	rows, err := p.pool.Query(p.ctx, `SELECT articles.id, articles.author_id,
articles.title, articles.content, articles.created_at, authors.name FROM articles, authors WHERE articles.author_id = authors.id;`)
	defer rows.Close()
	if err == pgx.ErrNoRows {
		return []storage.Article{}, nil

	}
	if err != nil {
		return []storage.Article{}, err
	}
	var result []storage.Article
	for rows.Next() {
		var p storage.Article
		err = rows.Scan(&p.ID, &p.AuthorID, &p.Title, &p.Content, &p.CreatedAt, &p.AuthorName)

		if err != nil {
			return []storage.Article{}, err
		}
		result = append(result, p)
	}
	err = rows.Err()
	if err != nil {
		return []storage.Article{}, err
	}
	return result, nil
}

func (p *postgresDB) AddArticle(article storage.Article) error { // создание новой публикации
	_, err := p.pool.Exec(p.ctx, `INSERT INTO articles (author_id,
title, content, created_at) VALUES ($1, $2, $3, $4);`,
		article.AuthorID, article.Title, article.Content, article.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresDB) UpdateArticle(article storage.Article) error { // обновление публикации
	_, err := p.pool.Exec(p.ctx, `UPDATE articles SET author_id = $1,
title = $2, content = $3, created_at = $4 WHERE articles.id = $5;`, article.AuthorID,
		article.Title, article.Content, article.CreatedAt, article.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *postgresDB) DeleteArticle(article storage.Article) error { // удаление публикации по ID
	_, err := p.pool.Exec(p.ctx, `DELETE FROM articles WHERE articles.id = $1;`, article.ID)
	if err != nil {
		return err
	}
	return nil
}

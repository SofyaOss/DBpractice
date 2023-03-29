package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"skillfactory/DBpractice/pkg/storage"
)

const (
	database   string = "articlesdb"
	collection string = "articles"
)

type mongoDB struct {
	ctx  context.Context
	pool *mongo.Client
}

func New() (*mongoDB, error) {
	mongoOpts := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return &mongoDB{ctx: context.Background(), pool: client}, nil
}

func (m *mongoDB) ArticlesList() ([]storage.Article, error) { // получение всех публикаций
	collection := m.pool.Database(database).Collection(collection)
	filter := bson.D{}
	cur, err := collection.Find(m.ctx, filter)
	defer cur.Close(m.ctx)
	if err != nil {
		return []storage.Article{}, err
	}
	var result []storage.Article
	for cur.Next(m.ctx) {
		var p storage.Article
		err = cur.Decode(&p)
		if err != nil {
			return []storage.Article{}, err
		}
		result = append(result, p)

	}
	err = cur.Err()
	if err != nil {
		return []storage.Article{}, err
	}

	return result, nil
}
func (m *mongoDB) AddArticle(article storage.Article) error { // создание новой публикации
	collection := m.pool.Database(database).Collection(collection)
	_, err := collection.InsertOne(m.ctx, article)
	if err != nil {
		return err
	}
	return nil
}
func (m *mongoDB) UpdateArticle(article storage.Article) error { // обновление публикации
	collection := m.pool.Database(database).Collection(collection)
	filter := bson.M{"id": article.ID}
	set := bson.D{{"$set", bson.D{{"title", article.Title}, {"content", article.Content},
		{"authorid", article.AuthorID}, {"authorname", article.AuthorName},
		{"createdat", article.CreatedAt}, {"publishedat", article.PublishedAt}}}}
	_, err := collection.UpdateOne(m.ctx, filter, set)
	if err != nil {
		return err
	}
	return nil
}
func (m *mongoDB) DeleteArticle(article storage.Article) error { // удаление публикации по ID
	collection := m.pool.Database(database).Collection(collection)
	filter := bson.M{"id": article.ID}
	_, err := collection.DeleteOne(m.ctx, filter)
	if err != nil {
		return err
	}
	return nil
}

func (m *mongoDB) Close() { // закрытие соединения, общее для интерфейса
	err := m.pool.Disconnect(m.ctx)
	if err != nil {
		log.Fatal(err)
	}
}

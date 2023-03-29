package main

import (
	"log"
	"net/http"
	"skillfactory/DBpractice/pkg/api"
	"skillfactory/DBpractice/pkg/storage"
	"skillfactory/DBpractice/pkg/storage/memdb"
	"skillfactory/DBpractice/pkg/storage/mongo"
	"skillfactory/DBpractice/pkg/storage/postgres"
)

type server struct {
	db  storage.DBInterface
	api *api.API
}

func main() {
	var srv server
	db := memdb.New()

	// PostgreSQL
	db2, err := postgres.New()
	if err != nil {
		log.Fatal(err)
	}

	// MongoDB
	db3, err := mongo.New()
	if err != nil {
		log.Fatal(err)
	}
	_, _, _ = db, db2, db3
	srv.db = db2
	srv.api = api.New(srv.db)

	http.ListenAndServe(":8080", srv.api.Router())
}

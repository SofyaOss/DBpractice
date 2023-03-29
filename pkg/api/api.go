package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"skillfactory/DBpractice/pkg/storage"
)

type API struct {
	db     storage.DBInterface
	router *mux.Router
}

func New(db storage.DBInterface) *API {
	api := API{
		db: db,
	}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

func (api *API) endpoints() {
	api.router.HandleFunc("/articles", api.articlesHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/articles", api.addArticleHandler).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/articles", api.updateArticleHandler).Methods(http.MethodPut, http.MethodOptions)
	api.router.HandleFunc("/articles", api.deleteArticleHandler).Methods(http.MethodDelete, http.MethodOptions)
}

func (api *API) Router() *mux.Router {
	return api.router
}

func (api *API) articlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := api.db.ArticlesList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func (api *API) addArticleHandler(w http.ResponseWriter, r *http.Request) {
	var p storage.Article
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.AddArticle(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (api *API) updateArticleHandler(w http.ResponseWriter, r *http.Request) {
	var p storage.Article
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.UpdateArticle(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (api *API) deleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	var p storage.Article
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.DeleteArticle(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

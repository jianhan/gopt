package handler

import (
	"github.com/allegro/bigcache"
	"github.com/gorilla/mux"
	"net/http"
)

type zomatoPlace struct {
	cache *bigcache.BigCache
}

func NewZomatoPlace(cache *bigcache.BigCache) APIRouter {
	return &zomatoPlace{cache: cache}
}

func (p *zomatoPlace) SetupSubrouter(parentRouter *mux.Router) {
	r := parentRouter.PathPrefix("/zomato").Subrouter().StrictSlash(true)
	r.HandleFunc("/search", p.search).Name("get.zomato.search").Methods("GET")
	r.HandleFunc("/collections", p.collections).Name("get.zomato.collections").Methods("GET")
	r.HandleFunc("/categories", p.categories).Name("get.zomato.categories").Methods("GET")
}

func (p *zomatoPlace) search(rsp http.ResponseWriter, req *http.Request) {

}

func (p *zomatoPlace) collections(rsp http.ResponseWriter, req *http.Request) {

}

func (p *zomatoPlace) categories(rsp http.ResponseWriter, req *http.Request) {

}

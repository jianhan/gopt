package handler

import (
	"github.com/allegro/bigcache"
	"github.com/gorilla/mux"
	"net/http"
)

type zomato struct {
	cache *bigcache.BigCache
}

func NewZomatoPlace(cache *bigcache.BigCache) APIRouter {
	return &zomato{cache: cache}
}

func (z *zomato) SetupSubrouter(parentRouter *mux.Router) {
	r := parentRouter.PathPrefix("/zomato").Subrouter().StrictSlash(true)
	r.HandleFunc("/search", z.search).Name("get.zomato.search").Methods("GET")
	r.HandleFunc("/collections", z.collections).Name("get.zomato.collections").Methods("GET")
	r.HandleFunc("/categories", z.categories).Name("get.zomato.categories").Methods("GET")
	r.HandleFunc("/reviews", z.reviews).Name("get.zomato.reviews").Methods("GET")
	r.HandleFunc("/restaurant", z.restaurant).Name("get.zomato.restaurant").Methods("GET")
	r.HandleFunc("/daily-menu", z.dailyMenu).Name("get.zomato.daily-menu").Methods("GET")

}

func (z *zomato) search(rsp http.ResponseWriter, req *http.Request) {

}

func (z *zomato) collections(rsp http.ResponseWriter, req *http.Request) {

}

func (z *zomato) categories(rsp http.ResponseWriter, req *http.Request) {

}

func (z *zomato) reviews(rsp http.ResponseWriter, req *http.Request) {

}

func (z *zomato) restaurant(rsp http.ResponseWriter, req *http.Request) {

}

func (z *zomato) dailyMenu(rsp http.ResponseWriter, req *http.Request) {

}

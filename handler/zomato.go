package handler

import (
	"encoding/json"
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/gorilla/mux"
	ghttp "github.com/jianhan/gopt/http"
	gzomato "github.com/jianhan/gopt/zomato"
	"net/http"
)

type zomato struct {
	cache     *bigcache.BigCache
	commonAPI gzomato.CommonAPI
}

func NewZomatoPlace(cache *bigcache.BigCache, commonAPI gzomato.CommonAPI) APIRouter {
	return &zomato{cache: cache, commonAPI: commonAPI}
}

func (z *zomato) SetupSubrouter(parentRouter *mux.Router) {
	r := parentRouter.PathPrefix("/zomato").Subrouter().StrictSlash(true)
	r.HandleFunc("/search", z.search).Name(fmt.Sprintf("get.zomato.search")).Methods("GET")
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
	cacheKey := "get.zomato.categories"
	cachedResponse, cErr := z.cache.Get(cacheKey)
	if cErr != nil {
		categories, err := z.commonAPI.Categories()
		if err != nil {
			ghttp.SendJSONResponse(
				rsp,
				http.StatusInternalServerError,
				ghttp.HttpError{Message: err.Error(), Status: http.StatusInternalServerError},
			)
			return
		}

		jsonStrRsp, jErr := json.Marshal(&categories)
		if jErr != nil {
			ghttp.SendJSONResponse(
				rsp,
				http.StatusInternalServerError,
				ghttp.HttpError{Message: fmt.Sprintf("system error, unable to marshal request, %s", jErr.Error()), Status: http.StatusInternalServerError},
			)
			return
		}
		z.cache.Set(cacheKey, jsonStrRsp)
	}

	rsp.Write(cachedResponse)
	return

}

func (z *zomato) reviews(rsp http.ResponseWriter, req *http.Request) {

}

func (z *zomato) restaurant(rsp http.ResponseWriter, req *http.Request) {

}

func (z *zomato) dailyMenu(rsp http.ResponseWriter, req *http.Request) {

}

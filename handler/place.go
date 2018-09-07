package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

type place struct {
}

func NewPlace() APIRouter {
	return &user{}
}

func (p *place) SetupSubrouter(parentRouter *mux.Router) {
	r := parentRouter.PathPrefix("/place").Subrouter().StrictSlash(true)
	r.HandleFunc("/search", p.search).Name("get.place.search").Methods("GET")
}

func (p *place) search(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("tset SEARCH"))
}

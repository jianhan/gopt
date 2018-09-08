package handler

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
	"net/http"
)

type place struct {
	client *maps.Client
}

func NewPlace(client *maps.Client) APIRouter {
	return &place{client: client}
}

func (p *place) SetupSubrouter(parentRouter *mux.Router) {
	r := parentRouter.PathPrefix("/place").Subrouter().StrictSlash(true)
	r.HandleFunc("/search", p.search).Name("get.place.search").Methods("GET")
}

func (p *place) search(rsp http.ResponseWriter, req *http.Request) {
	findPlaceReq := &maps.FindPlaceFromTextRequest{
		Input:     "Museum of Contemporary Art Australia",
		InputType: maps.FindPlaceFromTextInputTypeTextQuery,
	}
	prsp, _ := p.client.FindPlaceFromText(req.Context(), findPlaceReq)
	for k := range prsp.Candidates {
		logrus.Info(prsp.Candidates[k].Name, "***********")
	}
	rsp.Write([]byte("tset SEARCH"))
}

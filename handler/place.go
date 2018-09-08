package handler

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	ghttp "github.com/jianhan/gopt/http"
	gplace "github.com/jianhan/gopt/place"
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

type SearchRequest struct {
	PriceLevel string `schema:"price_level"`
	Radius     uint
	Location   string
	Keyword    string
	Language   string
	MinPrice   string `schema:"min_price"`
	MaxPrice   string `schema:"max_price"`
	Name       string
	OpenNow    bool   `schema:"open_now"`
	RankBy     string `schema:"rank_by"`
	PlaceType  string `schema:"type"`
	PageToken  string `schema: "page_token"`
}

func (p *place) search(rsp http.ResponseWriter, req *http.Request) {
	r, rErr := gplace.NewNearbySearchRequest(
		gplace.NearbySearchRequestOptions{}.Location("-27.470125,153.021072"),
		gplace.NearbySearchRequestOptions{}.Raidus(1000),
	)

	searchRequest := new(SearchRequest)
	decoder := schema.NewDecoder()
	decoder.Decode(searchRequest, req.URL.Query())
	logrus.Info(searchRequest)
	return
	if rErr != nil {
		ghttp.SendJSONResponse(rsp, http.StatusInternalServerError, rErr)
		return
	}

	sRsp, sErr := p.client.NearbySearch(req.Context(), r)
	if sErr != nil {
		ghttp.SendJSONResponse(
			rsp,
			http.StatusInternalServerError,
			ghttp.HttpError{Message: fmt.Sprintf("unable fetch search results, %s", sErr.Error()), Status: http.StatusBadRequest},
		)
		return
	}

	logrus.Info(len(sRsp.Results))
}

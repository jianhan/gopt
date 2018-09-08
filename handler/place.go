package handler

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	ghttp "github.com/jianhan/gopt/http"
	gplace "github.com/jianhan/gopt/place"
	"github.com/leebenson/conform"
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
	Name      string `conform:"trim" valid:"required~Name is required"`
	Radius    uint
	Location  string `conform:"trim" valid:"required~Location is required"`
	Keyword   string `conform:"trim"`
	Language  string `conform:"trim"`
	MinPrice  string `schema:"min_price" conform:"trim"`
	MaxPrice  string `schema:"max_price" conform:"trim"`
	OpenNow   bool   `schema:"open_now"`
	RankBy    string `schema:"rank_by" conform:"trim"`
	PlaceType string `schema:"type" conform:"trim"`
	PageToken string `schema:"page_token" conform:"trim"`
}

func (s *SearchRequest) GenerateNearBySearchRequestOptions() ([]gplace.NearbySearchRequestOption) {
	options := []gplace.NearbySearchRequestOption{}
	if s.Name != "" {
		options = append(options, gplace.NearbySearchRequestOptions{}.Name(s.Name))
	}

	if s.Radius > 0 {
		options = append(options, gplace.NearbySearchRequestOptions{}.Raidus(s.Radius))
	}

	if s.Location != "" {
		options = append(options, gplace.NearbySearchRequestOptions{}.Location(s.Location))
	}

	if s.Keyword != "" {
		options = append(options, gplace.NearbySearchRequestOptions{}.Location(s.Keyword))
	}

	if s.Language != "" {
		options = append(options, gplace.NearbySearchRequestOptions{}.Language(s.Language))
	}

	if s.MinPrice != "" {
		options = append(options, gplace.NearbySearchRequestOptions{}.MinPrice(s.MinPrice))
	}

	if s.MaxPrice != "" {
		options = append(options, gplace.NearbySearchRequestOptions{}.MaxPrice(s.MaxPrice))
	}

	options = append(options, gplace.NearbySearchRequestOptions{}.OpenNow(s.OpenNow))

	if s.RankBy != "" {
		options = append(options, gplace.NearbySearchRequestOptions{}.RankBy(s.RankBy))
	}

	if s.PlaceType != "" {
		options = append(options, gplace.NearbySearchRequestOptions{}.Type(s.PlaceType))
	}

	if s.PageToken != "" {
		options = append(options, gplace.NearbySearchRequestOptions{}.PageToken(s.PageToken))
	}

	return options
}

func (p *place) search(rsp http.ResponseWriter, req *http.Request) {

	searchRequest := new(SearchRequest)
	decoder := schema.NewDecoder()
	decoder.Decode(searchRequest, req.URL.Query())
	conform.Strings(&searchRequest)

	logrus.Info(len(searchRequest.GenerateNearBySearchRequestOptions()))
	if _, vErr := govalidator.ValidateStruct(searchRequest); vErr != nil {
		ghttp.SendJSONResponse(
			rsp,
			http.StatusInternalServerError,
			ghttp.HttpError{Message: vErr.Error(), Status: http.StatusBadRequest},
		)
		return
		logrus.Info(vErr)
	}
	r, rErr := gplace.NewNearbySearchRequest(
		gplace.NearbySearchRequestOptions{}.Location("-27.470125,153.021072"),
		gplace.NearbySearchRequestOptions{}.Raidus(1000),
	)
	return
	if rErr != nil {
		ghttp.SendJSONResponse(
			rsp,
			http.StatusInternalServerError,
			ghttp.HttpError{Message: fmt.Sprintf("unable fetch search results, %s", rErr.Error()), Status: http.StatusInternalServerError},
		)
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

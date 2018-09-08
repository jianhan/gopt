package handler

import (
	"encoding/json"
	"fmt"
	"github.com/allegro/bigcache"
	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	ghttp "github.com/jianhan/gopt/http"
	gplace "github.com/jianhan/gopt/place"
	"github.com/leebenson/conform"
	"googlemaps.github.io/maps"
	"net/http"
)

type googlePlace struct {
	client *maps.Client
	cache  *bigcache.BigCache
}

func NewGooglePlace(client *maps.Client, cache *bigcache.BigCache) APIRouter {
	return &googlePlace{client: client, cache: cache}
}

func (p *googlePlace) SetupSubrouter(parentRouter *mux.Router) {
	r := parentRouter.PathPrefix("/google").Subrouter().StrictSlash(true)
	r.HandleFunc("/nearby-search", p.nearbySearch).Name("get.place.nearby-search").Methods("GET")
}

type GoogleSearchRequest struct {
	Name      string `conform:"trim" json:"name"`
	Radius    uint   `json:"radius"`
	Location  string `conform:"trim" valid:"required~Location is required" json:"location"`
	Keyword   string `conform:"trim" json:"keyword"`
	Language  string `conform:"trim" json:"language"`
	MinPrice  string `schema:"min_price" conform:"trim" json:"min_price"`
	MaxPrice  string `schema:"max_price" conform:"trim" json:"max_price"`
	OpenNow   bool   `schema:"open_now" json:"open_now"`
	RankBy    string `schema:"rank_by" conform:"trim" json:"rank_by"`
	PlaceType string `schema:"type" conform:"trim" json:"price_type"`
	PageToken string `schema:"page_token" conform:"trim" json:"page_token"`
}

func (s *GoogleSearchRequest) GenerateNearBySearchRequestOptions() ([]gplace.NearbySearchRequestOption) {
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
		options = append(options, gplace.NearbySearchRequestOptions{}.Keyword(s.Keyword))
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

func (p *googlePlace) nearbySearch(rsp http.ResponseWriter, req *http.Request) {
	searchRequest := new(GoogleSearchRequest)
	schema.NewDecoder().Decode(searchRequest, req.URL.Query())
	conform.Strings(&searchRequest)

	jsonStr, jErr := json.Marshal(&searchRequest)
	if jErr != nil {
		ghttp.SendJSONResponse(
			rsp,
			http.StatusInternalServerError,
			ghttp.HttpError{Message: fmt.Sprintf("system error, unable to marshal request, %s", jErr.Error()), Status: http.StatusInternalServerError},
		)
		return
	}
	cacheKey := string(jsonStr)

	cachedResponse, cErr := p.cache.Get(cacheKey)
	if cErr != nil {
		// no entry exist in cache
		if _, vErr := govalidator.ValidateStruct(searchRequest); vErr != nil {
			ghttp.SendJSONResponse(
				rsp,
				http.StatusInternalServerError,
				ghttp.HttpError{Message: vErr.Error(), Status: http.StatusInternalServerError},
			)
			return
		}

		nsReq, rErr := gplace.NewNearbySearchRequest(searchRequest.GenerateNearBySearchRequestOptions()...)
		if rErr != nil {
			ghttp.SendJSONResponse(
				rsp,
				http.StatusInternalServerError,
				ghttp.HttpError{Message: fmt.Sprintf("unable fetch search results, %s", rErr.Error()), Status: http.StatusInternalServerError},
			)
			return
		}

		sRsp, sErr := p.client.NearbySearch(req.Context(), nsReq)
		if sErr != nil {
			ghttp.SendJSONResponse(
				rsp,
				http.StatusInternalServerError,
				ghttp.HttpError{Message: fmt.Sprintf("unable fetch search results, %s", sErr.Error()), Status: http.StatusBadRequest},
			)
			return
		}

		jsonStrRsp, jErr := json.Marshal(&sRsp)
		if jErr != nil {
			ghttp.SendJSONResponse(
				rsp,
				http.StatusInternalServerError,
				ghttp.HttpError{Message: fmt.Sprintf("system error, unable to marshal request, %s", jErr.Error()), Status: http.StatusInternalServerError},
			)
			return
		}
		p.cache.Set(cacheKey, jsonStrRsp)

		ghttp.SendJSONResponse(
			rsp,
			http.StatusOK,
			sRsp,
		)
		return
	}

	rsp.Write(cachedResponse)
	return
}

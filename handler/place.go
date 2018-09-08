package handler

import (
	"fmt"
	"github.com/gorilla/mux"
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

func (p *place) search(rsp http.ResponseWriter, req *http.Request) {
	//findPlaceReq := &maps.FindPlaceFromTextRequest{
	//	Input:     "Museum of Contemporary Art Australia",
	//	InputType: maps.FindPlaceFromTextInputTypeTextQuery,
	//}
	//prsp, _ := p.client.FindPlaceFromText(req.Context(), findPlaceReq)
	//for k := range prsp.Candidates {
	//	logrus.Info(prsp.Candidates[k].Name, "***********")
	//}

	//resp, err := place.Get
	//NearbySearch(context.Background(), r)
	client, err := gplace.GetClient()
	if err != nil {
		ghttp.SendJSONResponse(rsp, http.StatusInternalServerError, ghttp.HttpError{Message: "unable to get google place client", Status: http.StatusInternalServerError})
		return
	}

	r, rErr := gplace.NewNearbySearchRequest(
		gplace.NearbySearchRequestOptions{}.Location("-27.470125,153.021072"),
		gplace.NearbySearchRequestOptions{}.Raidus(1000),
	)
	if rErr != nil {
		ghttp.SendJSONResponse(rsp, http.StatusInternalServerError, rErr)
		logrus.Error(rErr)
		return
	}

	sRsp, sErr := client.NearbySearch(req.Context(), r)
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

package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
)

type APIRouter interface {
	SetupSubrouter(*mux.Router)
}

type Router interface {
	GetRouter() (*mux.Router, error)
}

func NewRouter(middlewares []negroni.Handler, apiRouters []APIRouter) (Router, error) {
	r := &router{
		router:           mux.NewRouter().StrictSlash(true),
		commonMiddleware: middlewares,
		apiRouters:       apiRouters,
	}

	return r.initAPIRoutes().initWebRoutes(), nil
}

type router struct {
	router           *mux.Router
	commonMiddleware []negroni.Handler
	apiRouters       []APIRouter
}

func (r *router) GetRouter() (*mux.Router, error) {
	return r.router, nil
}

func (r *router) initAPIRoutes() *router {
	apiVersionSubrouter := mux.NewRouter().StrictSlash(true).PathPrefix("/api/v1").Subrouter()

	// setup sub routes
	for k := range r.apiRouters {
		r.apiRouters[k].SetupSubrouter(apiVersionSubrouter)
	}
	r.router.PathPrefix("/api/v1").Handler(negroni.New(r.commonMiddleware...).With(
		negroni.Wrap(apiVersionSubrouter),
	))

	return r
}

func (r *router) initWebRoutes() *router {
	return r
}

func SendJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	body, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
	logrus.Info(data)
	if _, err = w.Write(body); err != nil {
		// TODO: log
	}
}

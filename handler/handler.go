package handler

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

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

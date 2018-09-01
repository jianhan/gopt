package handler

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

type Router interface {
	GetRouter() (*mux.Router, error)
}

func NewRouter() (Router, error) {
	negroni.Classic()
	r := &router{
		router:           mux.NewRouter().StrictSlash(true),
		apiRoutes:        mux.NewRouter().StrictSlash(true),
		webRoutes:        mux.NewRouter().StrictSlash(true),
		commonMiddleware: []negroni.Handler{negroni.NewRecovery(), negroni.NewLogger()},
	}
	return r.initAPIRoutes().initWebRoutes(), nil
}

type router struct {
	router           *mux.Router
	apiRoutes        *mux.Router
	webRoutes        *mux.Router
	commonMiddleware []negroni.Handler
}

func (r *router) GetRouter() (*mux.Router, error) {
	return r.router, nil
}

func (r *router) initAPIRoutes() *router {

	// user sub router
	u := &user{}
	apiVersionSubrouter := mux.NewRouter().StrictSlash(true).PathPrefix("/api/v1").Subrouter()

	userSubrouter := apiVersionSubrouter.PathPrefix("/user").Subrouter()
	userSubrouter.HandleFunc("/profile", u.profile).Name("get.user.profile").Methods("GET")

	////r.router.HandleFunc("/profile", u.profile)
	//userRouter := mux.NewRouter().PathPrefix("/api/v1").Subrouter().StrictSlash(true)
	//userRouter.Handle("/", userSubrouter)

	r.router.PathPrefix("/api/v1").Handler(negroni.New(r.commonMiddleware...).With(
		negroni.Wrap(apiVersionSubrouter),
	))

	return r
}

func (r *router) initWebRoutes() *router {
	return r
}

package handler

import (
	"github.com/gorilla/mux"
)

type Router interface {
	GetRouter() (*mux.Router, error)
}

func NewRouter() (Router, error) {
	r := &router{
		router: mux.NewRouter(),
	}
	r.init()

	return r, nil
}

type router struct {
	router *mux.Router
}

func (r *router) GetRouter() (*mux.Router, error) {
	return r.router, nil
}

func (r *router) init() error {
	u := &user{}
	//apiSubrouter := r.router.PathPrefix("/api/v1").Subrouter()
	//usersSubrouter := apiSubrouter.PathPrefix("/user").Subrouter()
	r.router.HandleFunc("/user/profile", u.profile).Name("get.user.profile").Methods("GET")
	return nil
}

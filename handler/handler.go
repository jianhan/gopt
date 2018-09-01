package handler

import "github.com/gorilla/mux"

type Router interface {
	GetRouter() (*mux.Router, error)
}

func NewRouter() (Router, error) {
	r := &router{
		router: mux.NewRouter(),
	}

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
	r.router.HandleFunc("/", u.profile)

	return nil
}

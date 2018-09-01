package handler

import (
	"github.com/gorilla/mux"
	"net/http"
)

type user struct {
}

func NewUser() APIRouter {
	return &user{}
}

func (u *user) SetupSubrouter(parentRouter *mux.Router) {
	r := parentRouter.PathPrefix("/user").Subrouter().StrictSlash(true)
	r.HandleFunc("/profile", u.profile).Name("get.user.profile").Methods("GET")
	r.Use(checkAuth)
}

func (u *user) profile(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("tset"))
}

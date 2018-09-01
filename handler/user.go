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
	userSubrouter := parentRouter.PathPrefix("/user").Subrouter()
	userSubrouter.HandleFunc("/profile", u.profile).Name("get.user.profile").Methods("GET")
}

func (u *user) profile(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("tset"))
}

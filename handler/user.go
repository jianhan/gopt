package handler

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"net/http"
)

type user struct {
}

func NewUser() APIRouter {
	return &user{}
}

func (u *user) SetupSubrouter(parentRouter *mux.Router) {
	userSubrouter := mux.NewRouter().PathPrefix("/user").Subrouter().StrictSlash(true)
	userSubrouter.HandleFunc("/profile", u.profile).Name("get.user.profile").Methods("GET")
	parentRouter.PathPrefix("/user").Handler(negroni.New(
		negroni.HandlerFunc(CheckAuth),
		negroni.Wrap(userSubrouter),
	))
}

func (u *user) profile(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("tset"))
}

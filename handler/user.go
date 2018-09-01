package handler

import (
	"net/http"
)

type user struct {
}

func (u *user) profile(rsp http.ResponseWriter, req *http.Request) {
	rsp.Write([]byte("tset"))
}

package middleware

import "net/http"

func JSONRspHeader(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	rw.Header().Set("Content-Type", "application/json")
	next(rw, r)
}

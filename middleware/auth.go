package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/jianhan/gopt/firebase"
	ghttp "github.com/jianhan/gopt/http"
	"github.com/sirupsen/logrus"
)

func CheckAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get Authorization header
		idToken := r.Header.Get("Authorization")
		if idToken == "" {
			ghttp.SendJSONResponse(w, http.StatusUnauthorized, ghttp.HttpError{Status: http.StatusUnauthorized, Message: "Unauthorized"})
			return
		}

		// get token
		splitToken := strings.Split(idToken, "Bearer")
		idToken = strings.Trim(splitToken[1], " ")
		if idToken == "" {
			ghttp.SendJSONResponse(w, http.StatusUnauthorized, ghttp.HttpError{Status: http.StatusUnauthorized, Message: "Unauthorized, id token is missing"})
			return
		}

		// get firebase app
		firebaseApp, err := firebase.FirebaseApp()
		if err != nil {
			ghttp.SendJSONResponse(w, http.StatusInternalServerError, ghttp.HttpError{Status: http.StatusInternalServerError, Message: "Internal server error, unable to authenticate user"})
			return
		}

		// validate user
		client, err := firebaseApp.Auth(context.Background())
		if err != nil {
			ghttp.SendJSONResponse(w, http.StatusUnauthorized, ghttp.HttpError{Status: http.StatusUnauthorized, Message: "Invalid id token"})
			return
		}

		// get user
		user, err := client.VerifyIDToken(r.Context(), idToken)
		if err != nil {
			logrus.Errorf("error verifying ID token, %v: %s", err, idToken)
			ghttp.SendJSONResponse(w, http.StatusInternalServerError, ghttp.HttpError{Status: http.StatusInternalServerError, Message: "Unable to verify token"})
			return
		}
		logrus.Debug(user)

		// passed authentication
		next.ServeHTTP(w, r)
	})
}

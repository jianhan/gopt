package handler

import (
	"context"
	"github.com/jianhan/gopt/firebase"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func CheckAuth(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// get Authorization header
	idToken := r.Header.Get("Authorization")
	if idToken == "" {
		SendJSONResponse(rw, http.StatusUnauthorized, HttpError{Status: http.StatusUnauthorized, Message: "Unauthorized"})
		return
	}

	// get token
	splitToken := strings.Split(idToken, "Bearer")
	idToken = strings.Trim(splitToken[1], " ")
	if idToken == "" {
		SendJSONResponse(rw, http.StatusUnauthorized, HttpError{Status: http.StatusUnauthorized, Message: "Unauthorized, id token is missing"})
		return
	}

	// get firebase app
	firebaseApp, err := firebase.FirebaseApp()
	if err != nil {
		SendJSONResponse(rw, http.StatusInternalServerError, HttpError{Status: http.StatusInternalServerError, Message: "Internal server error, unable to authenticate user"})
		return
	}

	// validate user
	client, err := firebaseApp.Auth(context.Background())
	if err != nil {
		SendJSONResponse(rw, http.StatusUnauthorized, HttpError{Status: http.StatusUnauthorized, Message: "Invalid id token"})
		return
	}

	token, err := client.VerifyIDToken(r.Context(), idToken)
	if err != nil {
		logrus.Errorf("error verifying ID token", err, idToken)
		SendJSONResponse(rw, http.StatusInternalServerError, HttpError{Status: http.StatusInternalServerError, Message: "Unable to verify token"})
		return
	}

	logrus.Info(token)

	// passed authentication
	next(rw, r)
}

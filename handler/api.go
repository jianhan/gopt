package handler

import "github.com/gorilla/mux"

type APIRouter interface {
	SetupSubrouter(parentRouter *mux.Router)
}

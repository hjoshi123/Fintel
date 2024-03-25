package server

import (
	"github.com/gorilla/mux"
	"github.com/hjoshi123/fintel/infra/router"
)

func Setup() *mux.Router {
	r := mux.NewRouter()
	router.Initialize(r)
	return r
}

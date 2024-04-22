package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hjoshi123/fintel/infra/api"
	controllers "github.com/hjoshi123/fintel/pkg/controllers/sentiment"
)

func Initialize(r *mux.Router) {
	v0Router := r.PathPrefix("/v0").Subrouter()

	sentimentStockController := controllers.NewSentiStockController()
	sentimentRouter := v0Router.PathPrefix("/sentiment").Subrouter()
	sentimentRouter.Handle(fmt.Sprintf("/%s/{id}", sentimentStockController.Path()), api.CustomHandler(sentimentStockController.Show)).Methods(http.MethodGet)
}

package api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hjoshi123/fintel/infra/util"
	"github.com/hjoshi123/fintel/pkg/models"
)

type Input struct {
	ID        string
	W         http.ResponseWriter
	R         *http.Request
	GetParams map[string][]string
	User      *models.User
}

type Output struct {
	Output interface{}
}

type CustomHandler func(ctx context.Context, input Input) (Output, error)

func (c CustomHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	input := Input{}
	input.W = w
	input.R = r
	input.ID = mux.Vars(r)["id"]
	input.GetParams = make(map[string][]string)

	user, _ := ctx.Value("user").(*models.User)
	input.User = user

	for k, v := range r.URL.Query() {
		input.GetParams[k] = v
	}

	output, err := c(ctx, input)
	if err != nil {
		util.Log.Error().Ctx(ctx).Err(err).Msg("failed to process request")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if output.Output != nil {
		b, err := json.Marshal(output.Output)
		if err != nil {
			util.Log.Error().Ctx(ctx).Err(err).Msg("failed to marshal response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}

	w.WriteHeader(http.StatusOK)
}

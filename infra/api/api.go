package api

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/jsonapi"
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
	IsJSON bool
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
		if output.IsJSON {
			w.Header().Set("Content-Type", "application/json")
			b, err := json.Marshal(output.Output)
			if err != nil {
				util.Log.Error().Ctx(ctx).Err(err).Msg("failed to marshal response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(b)
		} else {
			log.Printf("output: %+#v\n", output.Output)
			resp, err := jsonapi.Marshal(output.Output)
			if err != nil {
				util.Log.Error().Ctx(ctx).Err(err).Msg("failed to marshal response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var b bytes.Buffer
			encodedResponse := bufio.NewWriter(&b)
			//fmt.Print(p)
			if err := json.NewEncoder(encodedResponse).Encode(resp); err != nil {
				util.Log.Error().Ctx(ctx).Err(err).Msg("failed to encode response")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			encodedResponse.Flush()
			w.Header().Set("Content-Type", jsonapi.MediaType)
			w.Write(b.Bytes())
			w.WriteHeader(http.StatusOK)
		}
	}
}

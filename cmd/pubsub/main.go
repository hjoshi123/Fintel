package main

import (
	"context"

	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/pubsub"
	"github.com/hjoshi123/fintel/infra/util"
)

func main() {
	ctx := context.Background()
	config.Load(ctx, "../../pkl/config/config.pkl")

	util.Logger()

	err := pubsub.Consume(ctx)
	if err != nil {
		util.Log.Fatal().Err(err).Msg("Failed to consume message")
	}
}

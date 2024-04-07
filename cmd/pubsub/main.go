package main

import (
	"context"
	"os"

	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/constants"
	"github.com/hjoshi123/fintel/infra/pubsub"
	"github.com/hjoshi123/fintel/infra/util"
	stockHelpers "github.com/hjoshi123/fintel/pkg/helpers/pubsub"
)

func main() {
	ctx := context.Background()

	pklPath := os.Getenv("PKL_PATH")
	config.Load(ctx, pklPath)
	util.Logger()

	client := pubsub.NewKafkaPubSub()

	stockHelp := stockHelpers.NewStockHelpers()
	err := client.Subscribe(ctx, constants.StocksNewsCreateTopic, stockHelp.StockNewsCreate)
	if err != nil {
		util.Log.Fatal().Err(err).Msg("Failed to consume message")
	}

	err = client.Consume(ctx)
	if err != nil {
		util.Log.Fatal().Err(err).Msg("Failed to consume message")
	}
}

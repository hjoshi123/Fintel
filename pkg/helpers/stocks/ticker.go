package helpers

import (
	"context"
	"sync"

	"github.com/gocarina/gocsv"
	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/thirdparty/alphavantage"
	"github.com/hjoshi123/fintel/infra/util"
	datastore "github.com/hjoshi123/fintel/pkg/datastore/db"
	"github.com/hjoshi123/fintel/pkg/models"
	"github.com/spf13/cobra"
)

const (
	alphaVantageBaseURL = "https://www.alphavantage.co/query"
)

type Ticker struct {
	Symbol    string `csv:"symbol"`
	Name      string `csv:"name"`
	AssetType string `csv:"assetType"`
	Status    string `csv:"status"`
}

func GetLatestTickers(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	av, err := alphavantage.NewAVClient(ctx, nil, alphaVantageBaseURL, config.Spec.AlphaVantageApiKey)
	if err != nil {
		util.Log.Error().Ctx(ctx).Err(err).Msg("Failed to create new alpha vantage client")
		return err
	}

	params := map[string]string{
		"function": "LISTING_STATUS",
	}

	req, err := av.NewRequest(ctx, "GET", nil, params)
	if err != nil {
		util.Log.Error().Ctx(ctx).Err(err).Msg("Failed to create new request")
		return err
	}

	resp, err := av.Do(ctx, req, nil)
	if err != nil {
		util.Log.Error().Ctx(ctx).Err(err).Msg("Failed to do request")
		return err
	}

	defer resp.Body.Close()

	tickers := make([]*Ticker, 0)

	err = gocsv.Unmarshal(resp.Body, &tickers)
	if err != nil {
		util.Log.Error().Ctx(ctx).Err(err).Msg("Failed to unmarshal response")
		return err
	}

	var wg sync.WaitGroup
	for _, ticker := range tickers {
		wg.Add(1)
		go func(t *Ticker) {
			defer wg.Done()
			getAndSaveTicker(ctx, t)
		}(ticker)
	}

	wg.Wait()
	return nil
}

func getAndSaveTicker(ctx context.Context, ticker *Ticker) {
	tickerStore := datastore.NewTickerStore()
	t := new(models.Ticker)

	t.Ticker = ticker.Symbol
	t.Name = ticker.Name

	err := tickerStore.Save(ctx, t)
	if err != nil {
		util.Log.Error().Ctx(ctx).Err(err).Msg("Failed to save ticker")
	}

	util.Log.Info().Ctx(ctx).Msg("Saved ticker")
}

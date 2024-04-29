package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/util"
	datastore "github.com/hjoshi123/fintel/pkg/datastore/db"
	"github.com/hjoshi123/fintel/pkg/external/finhistory"
	"github.com/hjoshi123/fintel/pkg/models"
	"github.com/spf13/cobra"
)

const (
	baseURL = "https://financialmodelingprep.com/api/v3"
	path    = "/historical-price-full"
)

func HandlePriceUpdate(ctx context.Context, msg *models.Message) error {
	priceFromPubsub := new(models.PriceUpdate)
	if err := json.NewDecoder(strings.NewReader(msg.Data)).Decode(priceFromPubsub); err != nil {
		return err
	}

	// Sleeping here because it needs to wait for all the things to be inserted in db
	time.Sleep(time.Second * 30)

	fromDate, err := time.Parse("2006-01-02", priceFromPubsub.FromDate)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to parse from date")
		return err
	}

	toDate, err := time.Parse("2006-01-02", priceFromPubsub.ToDate)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to parse to date")
		return err
	}

	err = GetDailyPrice(ctx, priceFromPubsub.Ticker, fromDate, toDate)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to get daily price")
		return err
	}

	return nil
}

func GetDailyPriceForTasker(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	if len(args) != 3 {
		return fmt.Errorf("Usage: %s <ticker> <fromDate> <toDate>", cmd.Use)
	}

	ticker := args[0]
	fromDate := args[1]
	toDate := args[2]

	from, err := time.Parse("2006-01-02", fromDate)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to parse from date")
		return err
	}

	to, err := time.Parse("2006-01-02", toDate)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to parse to date")
		return err
	}

	err = GetDailyPrice(ctx, ticker, from, to)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to get daily price")
		return err
	}

	return nil
}

func GetDailyPrice(ctx context.Context, ticker string, from, to time.Time) error {
	finnHistoryClient, err := finhistory.NewClient(baseURL, config.Spec.FinHistoryApiKey, nil)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to create finhistory client")
		return err
	}

	params := make(map[string]string, 0)
	params["apikey"] = finnHistoryClient.BearerToken
	params["from"] = from.Format("2006-01-02")
	params["to"] = to.Format("2006-01-02")
	req, err := finnHistoryClient.Request(ctx, http.MethodGet, fmt.Sprintf("%s/%s", path, ticker), params, nil)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to create request")
		return err
	}

	historicalPrices := new(finhistory.HistoricalPrices)
	resp, err := finnHistoryClient.Do(req, historicalPrices)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to get historical prices")
		return err
	}

	if resp.StatusCode != http.StatusOK {
		util.Log.Error().Err(err).Any("response", resp).Msg("Failed to get historical prices")
		return err
	}

	stockSentimentStore := datastore.NewStockSentimentStore()
	stocks, err := stockSentimentStore.GetByTickerAndTime(ctx, ticker, from, to)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to get stock sentiments")
		return err
	}

	for _, stock := range stocks {
		for _, price := range historicalPrices.Historical {
			if price.Date == stock.CreatedAt.Time.Format("2006-01-02") {
				stock.Price = price.Close
				err = stockSentimentStore.Save(ctx, stock, stock.R.Source.Name)
				if err != nil {
					util.Log.Error().Err(err).Msg("Failed to save stock sentiment")
				}
				break
			}
		}
	}

	return nil
}

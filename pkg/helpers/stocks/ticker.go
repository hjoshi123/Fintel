package helpers

import (
	"context"

	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/thirdparty/alphavantage"
	"github.com/hjoshi123/fintel/infra/util"
)

const (
	alphaVantageBaseURL = "https://www.alphavantage.co/query"
)

func GetLatestTickers(ctx context.Context) error {
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
}

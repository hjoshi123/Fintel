package helpers

import (
	"context"
	"time"

	"github.com/hjoshi123/fintel/infra/constants"
	"github.com/hjoshi123/fintel/infra/util"
	datastore "github.com/hjoshi123/fintel/pkg/datastore/db"
	datastoreIface "github.com/hjoshi123/fintel/pkg/datastore/interface"
	"github.com/hjoshi123/fintel/pkg/models"
)

type SentimentHelpers struct {
	SentimentStore  datastoreIface.StockSentimentStore
	TopContentStore datastoreIface.TopContentStore
}

func NewSentimentHelpers() *SentimentHelpers {
	return &SentimentHelpers{
		SentimentStore:  datastore.NewStockSentimentStore(),
		TopContentStore: datastore.NewTopContentStore(),
	}
}

func (s *SentimentHelpers) GetSentimentForStock(ctx context.Context, ticker string, timeStart, timeEnd time.Time) (*models.StockResponse, error) {
	stockSentiments, err := s.SentimentStore.GetByTickerAndTime(ctx, ticker, timeStart, timeEnd)
	if err != nil {
		util.Log.Error().Err(err).Msg("error getting associated stock sentiment")
		return nil, err
	}

	topContent, err := s.TopContentStore.GetByTickerAndTime(ctx, ticker, timeStart, timeEnd)
	if err != nil {
		util.Log.Error().Err(err).Msg("error getting top content")
		return nil, err
	}

	stockResponse := new(models.StockResponse)
	stockResponse.Ticker = ticker

	stockResponse.NewsSentiment = make([]*models.Sentiment, 0)
	stockResponse.SocialSentiment = make([]*models.Sentiment, 0)

	newsCorrelationPrices := make([]float64, 0)
	socialCorrelationPrices := make([]float64, 0)

	newsCorrelationSenitment := make([]float64, 0)
	socialCorrelationSentiment := make([]float64, 0)

	for _, stockSentiment := range stockSentiments {
		sent := new(models.Sentiment)
		sent.DailyICI = stockSentiment.DailyIci
		sent.ID = stockSentiment.ID
		sent.Date = &stockSentiment.CreatedAt.Time
		sent.Volume = stockSentiment.Chatter

		switch stockSentiment.R.GetSource().Name {
		case constants.StockNewsSource:
			newsCorrelationPrices = append(newsCorrelationPrices, stockSentiment.Price)
			newsCorrelationSenitment = append(newsCorrelationSenitment, stockSentiment.DailyIci)
			stockResponse.NewsSentiment = append(stockResponse.NewsSentiment, sent)
		case constants.StockSocialSource:
			socialCorrelationPrices = append(socialCorrelationPrices, stockSentiment.Price)
			socialCorrelationSentiment = append(socialCorrelationSentiment, stockSentiment.DailyIci)
			stockResponse.SocialSentiment = append(stockResponse.SocialSentiment, sent)
		}
	}

	stockResponse.NewsCorrelation, err = util.Pearson(newsCorrelationPrices, newsCorrelationSenitment)
	if err != nil {
		util.Log.Error().Err(err).Msg("error calculating news correlation")
	}

	stockResponse.SocialCorrelation, err = util.Pearson(socialCorrelationPrices, socialCorrelationSentiment)
	if err != nil {
		util.Log.Error().Err(err).Msg("error calculating social correlation")
	}

	stockResponse.TopContentsNews = make([]*models.TopContentResponse, 0)
	stockResponse.TopContentsSocial = make([]*models.TopContentResponse, 0)
	for _, content := range topContent {
		topContentResponse := new(models.TopContentResponse)
		topContentResponse.ID = content.ID
		topContentResponse.URL = content.URL
		topContentResponse.PostedDate = &content.CreatedAt

		switch content.R.GetSource().Name {
		case constants.StockNewsSource:
			stockResponse.TopContentsNews = append(stockResponse.TopContentsNews, topContentResponse)
		case constants.StockSocialSource:
			stockResponse.TopContentsSocial = append(stockResponse.TopContentsSocial, topContentResponse)
		}
	}

	util.Log.Debug().Any("stock", stockResponse).Msgf("stock response %s", ticker)
	return stockResponse, nil
}

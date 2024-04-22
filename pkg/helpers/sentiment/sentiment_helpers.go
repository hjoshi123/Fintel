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

	effectiveSocialICI := 0.0
	effectiveNewsICI := 0.0
	for _, stockSentiment := range stockSentiments {
		sent := new(models.Sentiment)
		sent.DailyICI = stockSentiment.DailyIci
		sent.ID = stockSentiment.ID
		sent.Date = &stockSentiment.CreatedAt.Time
		sent.Volume = stockSentiment.Chatter

		if stockSentiment.R.GetSource().Name == constants.StockNewsSource {
			effectiveNewsICI += stockSentiment.DailyIci
			stockResponse.NewsSentiment = append(stockResponse.NewsSentiment, sent)
		} else if stockSentiment.R.GetSource().Name == constants.StockSocialSource {
			effectiveSocialICI += stockSentiment.DailyIci
			stockResponse.SocialSentiment = append(stockResponse.SocialSentiment, sent)
		}
	}

	stockResponse.EffectiveNewsICI = (effectiveNewsICI / float64(len(stockResponse.NewsSentiment)))
	stockResponse.EffectiveSocialICI = (effectiveSocialICI / float64(len(stockResponse.SocialSentiment)))

	stockResponse.TopContentsNews = make([]*models.TopContentResponse, 0)
	stockResponse.TopContentsSocial = make([]*models.TopContentResponse, 0)
	for _, content := range topContent {
		topContentResponse := new(models.TopContentResponse)
		topContentResponse.ID = content.ID
		topContentResponse.URL = content.URL
		topContentResponse.PostedDate = &content.CreatedAt

		if content.R.GetSource().Name == constants.StockNewsSource {
			stockResponse.TopContentsNews = append(stockResponse.TopContentsNews, topContentResponse)
		} else if content.R.GetSource().Name == constants.StockSocialSource {
			stockResponse.TopContentsSocial = append(stockResponse.TopContentsSocial, topContentResponse)
		}
	}

	util.Log.Debug().Any("stock", stockResponse).Msgf("stock response %s", ticker)
	return stockResponse, nil
}

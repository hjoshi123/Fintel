package helpers

import (
	"context"
	"encoding/json"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/hjoshi123/fintel/infra/constants"
	"github.com/hjoshi123/fintel/infra/pubsub"
	"github.com/hjoshi123/fintel/infra/util"
	datastore "github.com/hjoshi123/fintel/pkg/datastore/db"
	datastoreIface "github.com/hjoshi123/fintel/pkg/datastore/interface"
	"github.com/hjoshi123/fintel/pkg/models"
	infoModels "github.com/hjoshi123/fintel/pkg/models/info"
	"github.com/volatiletech/null/v8"
)

type StockHelpers struct {
	Pubsub          pubsub.PubSub
	SentimentStore  datastoreIface.StockSentimentStore
	TopContentStore datastoreIface.TopContentStore
}

func NewStockHelpers(ctx context.Context) *StockHelpers {
	return &StockHelpers{
		Pubsub:          pubsub.NewKafkaPubSub(ctx),
		SentimentStore:  datastore.NewStockSentimentStore(),
		TopContentStore: datastore.NewTopContentStore(),
	}
}

func (s *StockHelpers) StocksCreateGeneric(ctx context.Context, msg *models.Message) error {
	util.Log.Info().Str("data", msg.Data).Msg("stock create generic")
	return nil
}

func (s *StockHelpers) StockNewsCreate(ctx context.Context, msg *models.Message) error {
	alphaNews := new(AlphaVantageNewsResponse)
	if err := json.NewDecoder(strings.NewReader(msg.Data)).Decode(alphaNews); err != nil {
		return err
	}

	stockSentiment := new(models.StockSentiment)
	stockSentiment.Ticker = alphaNews.Ticker
	chatter, err := strconv.Atoi(alphaNews.Items)
	if err != nil {
		util.Log.Error().Err(err).Msg("error converting items to int")
	}

	stockSentiment.Chatter = chatter
	positiveCount, negativeCount := getPositiveAndNegativeCount(alphaNews.Ticker, alphaNews.Feed)

	stockSentimentInfo := new(infoModels.StockSentimentInfo)
	stockSentimentInfo.PositiveCount = positiveCount
	stockSentimentInfo.NegativeCount = negativeCount
	stockSentimentInfo.NeutralCount = chatter - (positiveCount + negativeCount)

	jsonInfoBytes, err := json.Marshal(stockSentimentInfo)
	if err != nil {
		util.Log.Error().Err(err).Msg("error marshalling stock sentiment info")
		return err
	}

	stockSentiment.Info = null.JSON{
		JSON:  jsonInfoBytes,
		Valid: true,
	}

	util.Log.Info().Int("positive", positiveCount).Int("negative", negativeCount).Msg("positive and negative count")
	stockSentiment.DailyIci = calculateDailyICI(positiveCount, negativeCount)

	parsedTimeCreate, err := time.Parse("20060102", strings.Split(alphaNews.Feed[0].TimePublished, "T")[0])
	if err != nil {
		util.Log.Error().Err(err).Msg("error parsing time")
	}

	stockSentiment.CreatedAt = null.NewTime(parsedTimeCreate, true)
	stockSentiment.UpdatedAt = null.NewTime(time.Now(), true)

	err = s.SentimentStore.Save(ctx, stockSentiment, constants.StockNewsSource)
	if err != nil {
		util.Log.Error().Err(err).Msg("error saving stock sentiment")
		return err
	}

	feeds := make([]StockFeed, 0)
	if len(alphaNews.Feed) > 10 {
		feeds = alphaNews.Feed[:10]
	} else {
		feeds = alphaNews.Feed
	}

	for _, feed := range feeds {
		topContent := new(models.TopContent)
		topContent.Ticker = alphaNews.Ticker
		topContent.URL = feed.URL

		topContentInfo := new(infoModels.TopContentInfo)
		topContentInfo.Title = feed.Title
		topContentInfo.Summary = feed.Summary

		jsonBytes, err := topContentInfo.Scan()
		if err != nil {
			util.Log.Error().Err(err).Msg("error scanning top content info")
		}

		topContent.Info = null.JSON{
			JSON:  jsonBytes,
			Valid: true,
		}

		parsedTime, err := time.Parse("20060102", strings.Split(feed.TimePublished, "T")[0])
		if err != nil {
			util.Log.Error().Err(err).Msg("error parsing time")
		}

		topContent.CreatedAt = parsedTime
		topContent.UpdatedAt = time.Now()

		err = s.TopContentStore.Save(ctx, topContent, constants.StockNewsSource)
		if err != nil {
			util.Log.Error().Err(err).Msg("error saving top content")
			continue
		}
	}

	return nil
}

func (s *StockHelpers) StockSocialMediaCreate(ctx context.Context, msg *models.Message) error {
	redditResponse := new(SocialSentiment)
	if err := json.NewDecoder(strings.NewReader(msg.Data)).Decode(redditResponse); err != nil {
		return err
	}

	stockSentiment := new(models.StockSentiment)
	stockSentiment.Ticker = redditResponse.Ticker
	stockSentiment.Chatter = len(redditResponse.Feed)

	stockSentimentInfo := new(infoModels.StockSentimentInfo)

	positiveCount, negativeCount, neutralCount := 0, 0, 0
	for _, post := range redditResponse.Feed {
		if post.OverallSentimentScore.Compound > 0.15 {
			positiveCount++
		} else if post.OverallSentimentScore.Compound < 0 {
			negativeCount++
		} else {
			neutralCount++
		}
	}

	stockSentimentInfo.PositiveCount = positiveCount
	stockSentimentInfo.NegativeCount = negativeCount
	stockSentimentInfo.NeutralCount = neutralCount

	jsonInfoBytes, err := json.Marshal(stockSentimentInfo)
	if err != nil {
		util.Log.Error().Err(err).Msg("error marshalling stock sentiment info")
		return err
	}

	stockSentiment.Info = null.JSON{
		JSON:  jsonInfoBytes,
		Valid: true,
	}

	util.Log.Info().Int("positive", positiveCount).Int("negative", negativeCount).Msg("positive and negative count")
	stockSentiment.DailyIci = calculateDailyICI(positiveCount, negativeCount)

	parsedTimeCreate, err := time.Parse("2006-01-02", strings.Split(redditResponse.Feed[0].PostTime, " ")[0])
	if err != nil {
		util.Log.Error().Err(err).Msg("error parsing time")
	}
	stockSentiment.CreatedAt = null.NewTime(parsedTimeCreate, true)
	stockSentiment.UpdatedAt = null.NewTime(time.Now(), true)

	err = s.SentimentStore.Save(ctx, stockSentiment, constants.StockSocialSource)
	if err != nil {
		util.Log.Error().Err(err).Msg("error saving stock social sentiment")
		return err
	}

	for _, post := range redditResponse.Feed {
		topContent := new(models.TopContent)
		topContent.Ticker = redditResponse.Ticker
		topContent.URL = post.PostURL

		topContentInfo := new(infoModels.TopContentInfo)
		topContentInfo.Title = post.PostTitle
		topContentInfo.Summary = post.Body

		jsonBytes, err := topContentInfo.Scan()
		if err != nil {
			util.Log.Error().Err(err).Msg("error scanning top content info")
		}

		topContent.Info = null.JSON{
			JSON:  jsonBytes,
			Valid: true,
		}

		parsedTime, err := time.Parse("2006-01-02", strings.Split(post.PostTime, " ")[0])
		if err != nil {
			util.Log.Error().Err(err).Msg("error parsing time")
		}

		topContent.CreatedAt = parsedTime
		topContent.UpdatedAt = time.Now()

		err = s.TopContentStore.Save(ctx, topContent, constants.StockSocialSource)
		if err != nil {
			util.Log.Error().Err(err).Msg("error saving top content")
			continue
		}
	}

	return nil
}

func getPositiveAndNegativeCount(ticker string, articles []StockFeed) (int, int) {
	var positiveCount, negativeCount int
	for _, article := range articles {
		for _, tickSentiment := range article.TickerSentiment {
			if tickSentiment.Ticker == ticker {
				sentScoreFloat, err := strconv.ParseFloat(tickSentiment.TickerSentimentScore, 64)
				if err != nil {
					util.Log.Error().Err(err).Msg("error converting sentiment score to float")
					continue
				}
				if sentScoreFloat > 0.15 {
					positiveCount++
				} else if sentScoreFloat < -0.15 {
					negativeCount++
				}
			}
		}
	}
	return positiveCount, negativeCount
}

func calculateDailyICI(numberOfPositive, numberOfNegative int) float64 {
	// Natural log of ((1+number of positive)/(1+number of negative))
	return math.Log((1 + float64(numberOfPositive)) / (1 + float64(numberOfNegative)))
}

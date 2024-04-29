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
	"github.com/volatiletech/null/v8"
)

type StockHelpers struct {
	Pubsub          pubsub.PubSub
	SentimentStore  datastoreIface.StockSentimentStore
	TopContentStore datastoreIface.TopContentStore
}

func NewStockHelpers() *StockHelpers {
	return &StockHelpers{
		Pubsub:          pubsub.NewKafkaPubSub(),
		SentimentStore:  datastore.NewStockSentimentStore(),
		TopContentStore: datastore.NewTopContentStore(),
	}
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

	stockSentimentInfo := new(models.StockSentimentInfo)
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

	for _, feed := range alphaNews.Feed[:10] {
		topContent := new(models.TopContent)
		topContent.Ticker = alphaNews.Ticker
		topContent.URL = feed.URL

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

	stockSentimentInfo := new(models.StockSentimentInfo)

	positiveCount, negativeCount, neutralCount := 0, 0, 0
	for _, post := range redditResponse.Feed {
		if post.OverallSentimentScore.Compound > 0.15 {
			positiveCount++
		} else if post.OverallSentimentScore.Compound < -0.15 {
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

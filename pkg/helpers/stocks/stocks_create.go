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
	positiveCount, negativeCount := getPositiveAndNegativeCount(alphaNews.Feed)
	stockSentiment.DailyIci = calculateDailyICI(positiveCount, negativeCount)

	stockSentiment.CreatedAt = null.NewTime(time.Now(), true)
	stockSentiment.UpdatedAt = null.NewTime(time.Now(), true)

	err = s.SentimentStore.Save(ctx, stockSentiment)
	if err != nil {
		util.Log.Error().Err(err).Msg("error saving stock sentiment")
		return err
	}

	for _, feed := range alphaNews.Feed[:10] {
		topContent := new(models.TopContent)
		topContent.Ticker = alphaNews.Ticker
		topContent.URL = feed.URL
		topContent.SRC = constants.StockNewsSource

		parsedTime, err := time.Parse(time.RFC3339, feed.TimePublished)
		if err != nil {
			util.Log.Error().Err(err).Msg("error parsing time")
		}

		topContent.CreatedAt = parsedTime
		topContent.UpdatedAt = time.Now()

		err = s.TopContentStore.Save(ctx, topContent)
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
	stockSentiment.Chatter = redditResponse.Items

	positiveCount, negativeCount := 0, 0
	for _, post := range redditResponse.Feed {
		if post.OverallSentimentScore.Compound > 0 {
			positiveCount++
		} else if post.OverallSentimentScore.Compound < 0 {
			negativeCount++
		}
	}

	stockSentiment.DailyIci = calculateDailyICI(positiveCount, negativeCount)
	stockSentiment.CreatedAt = null.NewTime(time.Now(), true)
	stockSentiment.UpdatedAt = null.NewTime(time.Now(), true)

	err := s.SentimentStore.Save(ctx, stockSentiment)
	if err != nil {
		util.Log.Error().Err(err).Msg("error saving stock social sentiment")
		return err
	}

	for _, post := range redditResponse.Feed {
		topContent := new(models.TopContent)
		topContent.Ticker = redditResponse.Ticker
		topContent.URL = post.PostURL
		topContent.SRC = constants.StockRedditSource

		parsedTime, err := time.Parse(time.RFC3339, post.PostTime)
		if err != nil {
			util.Log.Error().Err(err).Msg("error parsing time")
		}

		topContent.CreatedAt = parsedTime
		topContent.UpdatedAt = time.Now()

		err = s.TopContentStore.Save(ctx, topContent)
		if err != nil {
			util.Log.Error().Err(err).Msg("error saving top content")
			continue
		}
	}

	return nil
}

func getPositiveAndNegativeCount(articles []StockFeed) (int, int) {
	var positiveCount, negativeCount int
	for _, article := range articles {
		if article.OverallSentimentScore > 0.15 {
			positiveCount++
		} else if article.OverallSentimentScore < -0.15 {
			negativeCount++
		}
	}
	return positiveCount, negativeCount
}

func calculateDailyICI(numberOfPositive, numberOfNegative int) float64 {
	// Natural log of ((1+number of positive)/(1+number of negative))
	return math.Log((1 + float64(numberOfPositive)) / (1 + float64(numberOfNegative)))
}

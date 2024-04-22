package datastore

import (
	"context"
	"time"

	"github.com/hjoshi123/fintel/pkg/models"
)

type StockSentimentStore interface {
	Save(ctx context.Context, stockSentiment *models.StockSentiment, src string) error
	GetByTickerAndTime(ctx context.Context, ticker string, timeStart, timeEnd time.Time) ([]*models.StockSentiment, error)
}

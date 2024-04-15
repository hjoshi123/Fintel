package datastore

import (
	"context"

	"github.com/hjoshi123/fintel/pkg/models"
)

type StockSentimentStore interface {
	Save(ctx context.Context, stockSentiment *models.StockSentiment) error
}

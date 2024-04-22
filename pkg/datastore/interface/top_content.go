package datastore

import (
	"context"
	"time"

	"github.com/hjoshi123/fintel/pkg/models"
)

type TopContentStore interface {
	Save(ctx context.Context, topContent *models.TopContent, src string) error
	GetByTickerAndTime(ctx context.Context, ticker string, timeStart, timeEnd time.Time) ([]*models.TopContent, error)
}

package datastore

import (
	"context"

	"github.com/hjoshi123/fintel/pkg/models"
)

type TickerStore interface {
	Save(ctx context.Context, ticker *models.Ticker) error
}

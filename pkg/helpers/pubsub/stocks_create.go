package helpers

import (
	"context"

	"github.com/hjoshi123/fintel/infra/pubsub"
	"github.com/hjoshi123/fintel/pkg/models"
)

type StockHelpers struct {
	Pubsub pubsub.PubSub
}

func NewStockHelpers() *StockHelpers {
	return &StockHelpers{
		Pubsub: pubsub.NewKafkaPubSub(),
	}
}

func (s *StockHelpers) StockNewsCreate(ctx context.Context, msg *models.Message) error {

}

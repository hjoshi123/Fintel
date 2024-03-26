package datastore

import (
	"context"

	"github.com/hjoshi123/fintel/infra/database"
	"github.com/hjoshi123/fintel/infra/util"
	datastore "github.com/hjoshi123/fintel/pkg/datastore/interface"
	"github.com/hjoshi123/fintel/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type tickerStore struct{}

func NewTickerStore() datastore.TickerStore {
	ts := new(tickerStore)
	return ts
}

func (ts *tickerStore) Save(ctx context.Context, ticker *models.Ticker) error {
	db := database.Connect()

	err := ticker.Insert(ctx, db, boil.Infer())
	if err != nil {
		util.Log.Error().Ctx(ctx).Err(err).Msg("Failed to insert ticker")
		return err
	}

	return nil
}

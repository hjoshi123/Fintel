package datastore

import (
	"context"
	"time"

	"github.com/hjoshi123/fintel/infra/database"
	datastore "github.com/hjoshi123/fintel/pkg/datastore/interface"
	"github.com/hjoshi123/fintel/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type stockSentimentStore struct {
}

func NewStockSentimentStore() datastore.StockSentimentStore {
	return &stockSentimentStore{}
}

func (s *stockSentimentStore) Save(ctx context.Context, stockSentiment *models.StockSentiment, src string) error {
	db := database.Connect()

	source, err := models.Sources(models.SourceWhere.Name.EQ(src)).One(ctx, db)
	if err != nil {
		return err
	}

	stockSentiment.SourceID = source.ID
	err = stockSentiment.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (s *stockSentimentStore) GetByTickerAndTime(ctx context.Context, ticker string, timeStart, timeEnd time.Time) ([]*models.StockSentiment, error) {
	db := database.Connect()

	stockSentiments, err := models.StockSentiments(
		qm.Where("ticker = ?", ticker),
		qm.And("created_at >= ?", timeStart),
		qm.And("created_at <= ?", timeEnd),
		qm.Load(models.StockSentimentRels.Source),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	return stockSentiments, nil
}

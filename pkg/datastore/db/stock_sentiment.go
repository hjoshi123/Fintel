package datastore

import (
	"context"

	"github.com/hjoshi123/fintel/infra/database"
	datastore "github.com/hjoshi123/fintel/pkg/datastore/interface"
	"github.com/hjoshi123/fintel/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type stockSentimentStore struct {
}

func NewStockSentimentStore() datastore.StockSentimentStore {
	return &stockSentimentStore{}
}

func (s *stockSentimentStore) Save(ctx context.Context, stockSentiment *models.StockSentiment) error {
	db := database.Connect()

	err := stockSentiment.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

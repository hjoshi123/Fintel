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

type topContent struct{}

func NewTopContentStore() datastore.TopContentStore {
	return &topContent{}
}

func (t *topContent) Save(ctx context.Context, topContent *models.TopContent, src string) error {
	db := database.Connect()

	source, err := models.Sources(models.SourceWhere.Name.EQ(src)).One(ctx, db)
	if err != nil {
		return err
	}

	topContent.SourceID = source.ID

	err = topContent.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (t *topContent) GetByTickerAndTime(ctx context.Context, ticker string, timeStart, timeEnd time.Time) ([]*models.TopContent, error) {
	db := database.Connect()

	topContents, err := models.TopContents(
		models.TopContentWhere.Ticker.EQ(ticker),
		models.TopContentWhere.CreatedAt.GTE(timeStart),
		models.TopContentWhere.CreatedAt.LTE(timeEnd),
		qm.Load(models.TopContentRels.Source),
	).All(ctx, db)
	if err != nil {
		return nil, err
	}

	return topContents, nil
}

package datastore

import (
	"context"

	"github.com/hjoshi123/fintel/infra/database"
	datastore "github.com/hjoshi123/fintel/pkg/datastore/interface"
	"github.com/hjoshi123/fintel/pkg/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type topContent struct{}

func NewTopContentStore() datastore.TopContentStore {
	return &topContent{}
}

func (t *topContent) Save(ctx context.Context, topContent *models.TopContent) error {
	db := database.Connect()

	err := topContent.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

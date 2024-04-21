package datastore

import (
	"context"

	"github.com/hjoshi123/fintel/pkg/models"
)

type TopContentStore interface {
	Save(ctx context.Context, topContent *models.TopContent) error
}

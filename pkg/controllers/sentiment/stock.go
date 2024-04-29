package controllers

import (
	"context"
	"fmt"
	"time"

	"github.com/hjoshi123/fintel/infra/api"
	"github.com/hjoshi123/fintel/infra/constants"
	helpers "github.com/hjoshi123/fintel/pkg/helpers/sentiment"
)

type SentimentStockController struct {
	SentimentHelpers *helpers.SentimentHelpers
}

func NewSentiStockController() *SentimentStockController {
	return &SentimentStockController{
		SentimentHelpers: helpers.NewSentimentHelpers(),
	}
}

func (s *SentimentStockController) Path() string {
	return "stock"
}

func (s *SentimentStockController) Show(ctx context.Context, input api.Input) (api.Output, error) {
	stockTicker := input.ID
	if stockTicker == "" {
		return api.Output{}, fmt.Errorf("stock ticker is required")
	}

	startTimeString, ok := input.GetParams[string(constants.TimeStart)]
	if !ok {
		return api.Output{}, fmt.Errorf("start time is required")
	}

	endTimeString, ok := input.GetParams[string(constants.TimeEnd)]
	if !ok {
		return api.Output{}, fmt.Errorf("end time is required")
	}

	startTime, err := time.Parse("2006-01-02", startTimeString[0])
	if err != nil {
		return api.Output{}, fmt.Errorf("error parsing start time %v", err.Error())
	}

	endTime, err := time.Parse("2006-01-02", endTimeString[0])
	if err != nil {
		return api.Output{}, fmt.Errorf("error parsing end time %v", err.Error())
	}

	stockRes, err := s.SentimentHelpers.GetSentimentForStock(ctx, stockTicker, startTime, endTime)
	if err != nil {
		return api.Output{}, fmt.Errorf("error getting stock sentiment %v", err.Error())
	}

	return api.Output{
		Output: stockRes,
	}, nil
}

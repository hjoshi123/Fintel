package run

import (
	helpers "github.com/hjoshi123/fintel/pkg/helpers/stocks"
	"github.com/spf13/cobra"
)

var (
	price = &cobra.Command{
		Use:   "price",
		Short: "Get Prices",
		Long:  "Get Prices within time range for a ticker",
		RunE:  helpers.GetDailyPriceForTasker,
	}
)

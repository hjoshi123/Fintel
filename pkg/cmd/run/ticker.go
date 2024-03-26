package run

import (
	helpers "github.com/hjoshi123/fintel/pkg/helpers/stocks"
	"github.com/spf13/cobra"
)

var (
	ticker = &cobra.Command{
		Use:   "ticker",
		Short: "Run the tasker with ticker",
		Long:  "Run the tasker with ticker",
		RunE:  helpers.GetLatestTickers,
	}
)

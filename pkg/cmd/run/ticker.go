package run

import "github.com/spf13/cobra"

var (
	ticker = &cobra.Command{
		Use:   "ticker",
		Short: "Run the tasker with ticker",
		Long:  "Run the tasker with ticker",
	}
)

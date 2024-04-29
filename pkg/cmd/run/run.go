package run

import "github.com/spf13/cobra"

var (
	run = &cobra.Command{
		Use:   "run",
		Short: "Run the tasker with task name and args",
		Long:  "Run the tasker with task name and args",
	}
)

func GetRunCommand() *cobra.Command {
	return run
}

// Add subcommands of run here
func init() {
	run.AddCommand(ticker, price)
}

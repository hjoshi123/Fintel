package cmd

import (
	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/database"
	"github.com/hjoshi123/fintel/infra/util"
	"github.com/hjoshi123/fintel/pkg/cmd/run"
	"github.com/spf13/cobra"
)

var (
	finTasker = &cobra.Command{
		Use:   "fintel",
		Short: "Start the tasker",
		Long:  "Run tasks using ./fintel run",
	}
)

func Execute() error {
	return finTasker.Execute()
}

func Initialize() {
	cobra.OnInitialize(initConfig)
	finTasker.AddCommand(run.GetRunCommand())
}

func initConfig() {
	ctx := finTasker.Context()
	config.Load(ctx, "../../pkl/config/config.pkl")
	_ = util.Logger()

	_ = database.Connect()

	util.Log.Info().Msg("Connected to database and loaded config")
}

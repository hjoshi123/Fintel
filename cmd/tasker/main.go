package main

import (
	"os"

	"github.com/hjoshi123/fintel/infra/util"
	"github.com/hjoshi123/fintel/pkg/cmd"
)

func main() {
	cmd.Initialize()
	if err := cmd.Execute(); err != nil {
		util.Log.Fatal().Err(err).Msg("Failed to run command")
		os.Exit(1)
	}
}

package api

import (
	"fmt"
	"net/http"
	"os"

	"ariga.io/atlas-go-sdk/atlasexec"
	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/database"
	"github.com/hjoshi123/fintel/infra/server"
	"github.com/hjoshi123/fintel/infra/util"
	"github.com/rs/cors"
	"github.com/spf13/cobra"
)

var (
	cobraServer = &cobra.Command{
		Use:   "server",
		Short: "Start the api server",
		RunE:  RunApiServer,
		Long:  "Start the flare. This will start the flare. Supported arguments are --disable-migration and --migration-path",
	}

	disableMigration bool
	migrationPath    string
)

func Execute() error {
	logger := util.Logger()
	cobraServer.Flags().StringVarP(&migrationPath, "migration-path", "m", "", "Path to the migration files")
	cobraServer.Flags().BoolVar(&disableMigration, "disable-migration", false, "Disable migration")
	if err := cobraServer.Execute(); err != nil {
		logger.Fatal().Err(err).Msg("Failed to start server")
		return err
	}

	return nil
}

func RunApiServer(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()
	config.Load(ctx, "../../pkl/config/config.pkl")
	logger := util.Logger()

	_ = database.Connect()

	if !disableMigration {
		workdir, err := atlasexec.NewWorkingDir(
			atlasexec.WithMigrations(
				os.DirFS("./migrations"),
			),
		)
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to create working directory")
			return err
		}
		// atlasexec works on a temporary directory, so we need to close it
		defer workdir.Close()

		// Initialize the client.
		client, err := atlasexec.NewClient(workdir.Path(), "atlas")
		if err != nil {
			logger.Fatal().Err(err).Msg("failed to create working directory")
			return err
		}

		client.MigrateApply(ctx, &atlasexec.MigrateApplyParams{
			URL: fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
				config.Spec.DB.DbAuth.DbUser, config.Spec.DB.DbAuth.DbPassword, config.Spec.DB.DbHost, config.Spec.DB.DbPort, config.Spec.DB.DbName),
		})
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8000", "http://localhost:3000"},
		AllowCredentials: true,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
	})

	httpMux := server.Setup()

	logger.Info().Msg("server started")
	logger.Debug().Msgf("config %#+v", config.Spec)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Spec.Port), c.Handler(httpMux)); err != nil {
		logger.Fatal().Err(err).Msg("failed to serve http server")
	}

	return nil
}

package database

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/util"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var (
	db   *sql.DB
	once sync.Once
)

func Connect() *sql.DB {
	once.Do(func() {
		logger := util.Logger()
		localDB, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
			config.Spec.DBUser, config.Spec.DBPassword, config.Spec.DBHost, config.Spec.DBPort, config.Spec.DBName))
		if err != nil {
			logger.Panic().Err(err).Msg("failed to connect to database")
		}

		if db == nil {
			db = localDB
		}
	})

	boil.SetDB(db)
	return db
}

// Code generated from Pkl module `com.github.hjoshi123.fintel.pkl.AppConfig`. DO NOT EDIT.
package appconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"github.com/hjoshi123/fintel/infra/config/appconfig/loglevel"
	"github.com/hjoshi123/fintel/infra/config/dbconfig"
	"github.com/hjoshi123/fintel/infra/config/kafkaconfig"
)

type AppConfig struct {
	Environment string `pkl:"environment"`

	Version string `pkl:"version"`

	Port uint16 `pkl:"port"`

	DB *dbconfig.DBConfig `pkl:"DB"`

	Kafka *kafkaconfig.KafkaConfig `pkl:"kafka"`

	AlphaVantageApiKey string `pkl:"alphaVantageApiKey"`

	LogLevel loglevel.LogLevel `pkl:"logLevel"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a AppConfig
func LoadFromPath(ctx context.Context, path string) (ret *AppConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a AppConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*AppConfig, error) {
	var ret AppConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

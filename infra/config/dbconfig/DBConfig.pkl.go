// Code generated from Pkl module `com.github.hjoshi123.fintel.pkl.DBConfig`. DO NOT EDIT.
package dbconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type DBConfig struct {
	DbPort string `pkl:"dbPort"`

	DbName string `pkl:"dbName"`

	DbHost string `pkl:"dbHost"`

	DbAuth *DBAuth `pkl:"dbAuth"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a DBConfig
func LoadFromPath(ctx context.Context, path string) (ret *DBConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a DBConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*DBConfig, error) {
	var ret DBConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

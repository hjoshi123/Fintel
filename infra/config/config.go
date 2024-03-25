package config

import (
	"context"
	"sync"

	"github.com/hjoshi123/fintel/infra/config/appconfig"
)

var (
	once sync.Once
	Spec *appconfig.AppConfig
)

func Load(ctx context.Context, path string) *appconfig.AppConfig {
	once.Do(func() {
		var err error
		spec, err := appconfig.LoadFromPath(ctx, path)
		if err != nil {
			panic(err)
		}

		if Spec == nil {
			Spec = spec
		}
	})

	return Spec
}

func IsDevelopment() bool {
	return Spec.Environment == "development"
}

func IsProduction() bool {
	return Spec.Environment == "production"
}

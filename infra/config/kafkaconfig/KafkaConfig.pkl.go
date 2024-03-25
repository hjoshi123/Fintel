// Code generated from Pkl module `com.github.hjoshi123.fintel.pkl.KafkaConfig`. DO NOT EDIT.
package kafkaconfig

import (
	"context"

	"github.com/apple/pkl-go/pkl"
)

type KafkaConfig struct {
	ClientID string `pkl:"clientID"`

	Brokers string `pkl:"brokers"`

	Group string `pkl:"group"`

	Topic string `pkl:"topic"`

	SaslEnable bool `pkl:"saslEnable"`

	SaslMechanism string `pkl:"saslMechanism"`

	Auth *KafkaAuth `pkl:"auth"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a KafkaConfig
func LoadFromPath(ctx context.Context, path string) (ret *KafkaConfig, err error) {
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

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a KafkaConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*KafkaConfig, error) {
	var ret KafkaConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

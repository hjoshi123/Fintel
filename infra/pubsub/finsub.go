package pubsub

import (
	"context"
	"crypto/tls"
	"reflect"
	"runtime"

	"github.com/IBM/sarama"
	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/util"
	"github.com/hjoshi123/fintel/pkg/models"
)

func GetSaramaConfig() *sarama.Config {
	util.Log.Info().Any("config", config.Spec)
	sConfig := sarama.NewConfig()
	sConfig.ClientID = config.Spec.KafkaClientID

	sConfig.Net.MaxOpenRequests = 1

	sConfig.Net.SASL.Enable = true
	sConfig.Net.SASL.Handshake = true
	sConfig.Net.SASL.Mechanism = "PLAIN"
	sConfig.Net.SASL.User = config.Spec.KafkaUsername
	sConfig.Net.SASL.Password = config.Spec.KafkaPassword
	sConfig.Net.TLS.Enable = true
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         0,
	}
	sConfig.Net.TLS.Config = tlsConfig

	// Producer config
	sConfig.Producer.Idempotent = true
	sConfig.Producer.Return.Errors = true
	sConfig.Producer.Return.Successes = true
	sConfig.Producer.RequiredAcks = sarama.WaitForAll
	sConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner
	sConfig.Producer.Transaction.Retry.Backoff = 10
	sConfig.Producer.Transaction.ID = config.Spec.KafkaTxnID

	// Consumer config
	sConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}
	sConfig.Consumer.IsolationLevel = sarama.ReadCommitted

	return sConfig
}

type MessageHandler func(context.Context, *models.Message) error

func (fn MessageHandler) GetFunctionName() string {
	v := reflect.ValueOf(fn)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			return rf.Name()
		}
	}
	return v.String()
}

type PubSub interface {
	Publish(ctx context.Context, msg *models.Message) error
	Subscribe(ctx context.Context, topic string, fn ...MessageHandler) error
	Consume(ctx context.Context) error
}

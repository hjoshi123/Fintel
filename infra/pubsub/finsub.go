package pubsub

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/constants"
	"github.com/hjoshi123/fintel/infra/util"
)

type Consumer struct {
	ready chan bool
}

func Consume(ctx context.Context) error {
	if !config.Spec.Kafka.SaslEnable {
		util.Log.Err(errors.New("sasl enable is not set to true")).Msg("sasl should be set to true")
		return errors.New("sasl enable is not set to true")
	}

	if len(config.Spec.Kafka.Brokers) == 0 {
		util.Log.Err(errors.New("no brokers found")).Msg("no brokers found")
		return errors.New("no brokers found")
	}

	if len(config.Spec.Kafka.Topic) == 0 {
		util.Log.Err(errors.New("no topics found")).Msg("no topics found")
		return errors.New("no topics found")
	}

	if len(config.Spec.Kafka.Group) == 0 {
		util.Log.Err(errors.New("no group found")).Msg("no group found")
		return errors.New("no group found")
	}

	keepRunning := true
	sConfig := sarama.NewConfig()
	sConfig.ClientID = config.Spec.Kafka.ClientID

	sConfig.Net.SASL.Enable = true
	sConfig.Net.SASL.Handshake = true
	sConfig.Net.SASL.Mechanism = "PLAIN"
	sConfig.Net.SASL.User = config.Spec.Kafka.Auth.Username
	sConfig.Net.SASL.Password = config.Spec.Kafka.Auth.Password
	sConfig.Net.TLS.Enable = true
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
		ClientAuth:         0,
	}
	sConfig.Net.TLS.Config = tlsConfig

	sConfig.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{sarama.NewBalanceStrategyRoundRobin()}

	consumer := Consumer{
		ready: make(chan bool),
	}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(strings.Split(config.Spec.Kafka.Brokers, ","), config.Spec.Kafka.Group, sConfig)
	if err != nil {
		util.Log.Err(err).Msg("error creating consumer group")
	}

	consumptionIsPaused := false
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := client.Consume(ctx, strings.Split(config.Spec.Kafka.Topic, ","), &consumer); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				util.Log.Err(err).Msg("error from consumer")
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			consumer.ready = make(chan bool)
		}
	}()

	<-consumer.ready // Await till the consumer has been set up
	util.Log.Info().Msg("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-ctx.Done():
			util.Log.Info().Msg("terminating: via context")
			keepRunning = false
		case <-sigterm:
			util.Log.Info().Msg("terminating: via signal")
			keepRunning = false
		case <-sigusr1:
			toggleConsumptionFlow(client, &consumptionIsPaused)
		}
	}
	cancel()
	wg.Wait()
	if err = client.Close(); err != nil {
		util.Log.Panic().Err(err).Msg("error closing client")
	}

	return nil
}

func toggleConsumptionFlow(client sarama.ConsumerGroup, isPaused *bool) {
	if *isPaused {
		client.ResumeAll()
		util.Log.Info().Msg("Resuming consumption")
	} else {
		client.PauseAll()
		util.Log.Info().Msg("Pausing consumption")
	}

	*isPaused = !*isPaused
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	close(consumer.ready)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				util.Log.Info().Msg("claim messages channel closed")
				return nil
			}
			switch string(message.Topic) {
			case constants.StocksNewsCreateTopic:
				util.Log.Debug().Msgf("StocksNewsCreateTopic: %s", string(message.Value))
			}
			session.MarkMessage(message, fmt.Sprintf("Read on: %s", time.Now().String()))
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

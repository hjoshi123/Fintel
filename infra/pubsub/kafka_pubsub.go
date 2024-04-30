package pubsub

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/hjoshi123/fintel/infra/config"
	"github.com/hjoshi123/fintel/infra/util"
	"github.com/hjoshi123/fintel/pkg/models"
)

type producerProvider struct {
	transactionIdGenerator int32
	producersLock          sync.Mutex
	producers              []sarama.AsyncProducer
	producerProvider       func() sarama.AsyncProducer
}

type KafkaPubSub struct {
	producerProvider *producerProvider
	handler          sync.Map
	ready            atomic.Int32
	kfkClient        sarama.Client
	shutdownContext  context.Context
	shutdownFunc     context.CancelFunc
}

func (p *producerProvider) borrow() (producer sarama.AsyncProducer) {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	if len(p.producers) == 0 {
		for {
			producer = p.producerProvider()
			if producer != nil {
				util.Log.Info().Msgf("Producer borrowed %v", producer)
				return
			}
		}
	}

	index := len(p.producers) - 1
	producer = p.producers[index]
	p.producers = p.producers[:index]
	return
}

func (p *producerProvider) release(producer sarama.AsyncProducer) {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	// If released producer is erroneous close it and don't return it to the producer pool.
	if producer.TxnStatus()&sarama.ProducerTxnFlagInError != 0 {
		// Try to close it
		_ = producer.Close()
		return
	}
	p.producers = append(p.producers, producer)
}

func (p *producerProvider) clear() {
	p.producersLock.Lock()
	defer p.producersLock.Unlock()

	for _, producer := range p.producers {
		producer.Close()
	}
	p.producers = p.producers[:0]
}

func newProducerProvider(brokers []string, config *sarama.Config) *producerProvider {
	provider := &producerProvider{}
	provider.producerProvider = func() sarama.AsyncProducer {
		producerConfig := *config
		// Append transactionIdGenerator to current config.Producer.Transaction.ID to ensure transaction-id uniqueness.
		suffix := provider.transactionIdGenerator
		// Append transactionIdGenerator to current config.Producer.Transaction.ID to ensure transaction-id uniqueness.
		if config.Producer.Transaction.ID != "" {
			provider.transactionIdGenerator++
			config.Producer.Transaction.ID = config.Producer.Transaction.ID + "-" + fmt.Sprint(suffix)
		}
		producer, err := sarama.NewAsyncProducer(brokers, &producerConfig)
		if err != nil {
			util.Log.Error().Err(err).Msg("Failed to create producer")
			return nil
		}
		return producer
	}
	return provider
}

var (
	pubsubOnce sync.Once
)

func NewKafkaPubSub(ctx context.Context) PubSub {
	kf := new(KafkaPubSub)
	kf.handler = sync.Map{}
	kf.ready = atomic.Int32{}
	shutdownContext, shutdownFunc := context.WithCancel(ctx)
	kf.shutdownContext = shutdownContext
	kf.shutdownFunc = shutdownFunc
	kf.producerProvider = newProducerProvider(strings.Split(config.Spec.KafkaBrokers, ","), GetSaramaConfig())
	pubsubOnce.Do(func() {
		kfkClient, err := sarama.NewClient(strings.Split(config.Spec.KafkaBrokers, ","), GetSaramaConfig())
		if err != nil {
			util.Log.Error().Err(err).Msg("Failed to create kafka client")
			return
		}

		kf.kfkClient = kfkClient
	})
	return kf
}

func (kf *KafkaPubSub) Shutdown() {
	kf.shutdownFunc()
	kf.producerProvider.clear()
	if kf.kfkClient != nil {
		kf.kfkClient.Close()
	}
	util.Log.Info().Msg("KafkaPubSub shutdown complete.")
}

func (kp *KafkaPubSub) Publish(ctx context.Context, message *models.Message) error {
	producer := kp.producerProvider.borrow()
	defer kp.producerProvider.release(producer)

	err := producer.BeginTxn()
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to begin transaction")
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: message.Topic,
		Value: sarama.StringEncoder(message.Data),
	}

	producer.Input() <- msg

	if err := producer.CommitTxn(); err != nil {
		for {
			if producer.TxnStatus()&sarama.ProducerTxnFlagFatalError != 0 {
				// fatal error. need to recreate producer.
				util.Log.Error().Err(err).Msg("Fatal error occured. Recreating producer")
				kp.producerProvider.clear() // Clear erroneous producers
				kp.producerProvider = newProducerProvider(strings.Split(config.Spec.KafkaBrokers, ","), GetSaramaConfig())
				continue
			}

			// If producer is in abortable state, try to abort current transaction.
			if producer.TxnStatus()&sarama.ProducerTxnFlagAbortableError != 0 {
				err = producer.AbortTxn()
				if err != nil {
					// If an error occured just retry it.
					util.Log.Error().Err(err).Msg("Failed to abort transaction")
					continue
				}
				break
			}
			// if not you can retry
			err = producer.CommitTxn()
			if err != nil {
				util.Log.Error().Err(err).Msg("Failed to commit transaction")
				continue
			}
		}
		return err
	}

	util.Log.Info().Msg("Message sent successfully and committed")

	return nil
}

func (kf *KafkaPubSub) Subscribe(ctx context.Context, topic string, fn ...MessageHandler) error {
	_, err := regexp.Compile(topic)
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to compile topic")
		return err
	}
	fnSlice, loaded := kf.handler.Load(topic)
	var fns []MessageHandler
	if loaded {
		fns = fnSlice.([]MessageHandler)
	} else {
		fns = make([]MessageHandler, 0)
	}
	fns = append(fns, fn...)

	kf.handler.Store(topic, fns)
	return nil
}

func (kf *KafkaPubSub) Consume(ctx context.Context) error {
	keepRunning := true

	availableTopics, err := kf.kfkClient.Topics()
	if err != nil {
		util.Log.Error().Err(err).Msg("Failed to get topics")
		return err
	}

	actualTopics := make([]string, 0)
	for _, topic := range availableTopics {
		if !strings.HasPrefix(topic, "__") {
			actualTopics = append(actualTopics, topic)
		}
	}

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(ctx)

	// Let's say main topic is "event.create." and we have 2 subtopics in the consumer end "event.create.slack" and "event.create.google"
	consumerGroupsTopics := make(map[string]bool)
	for _, topic := range actualTopics {
		consumerGroupsTopics[topic] = false
	}

	consumerGroups := util.GetSubTopicsFromTopics(consumerGroupsTopics, ".")

	for _, subTopics := range consumerGroups {
		wg.Add(1)
		util.Log.Info().Any("subtopics", subTopics).Msg("Starting consumer group")
		go kf.startConsumerGroupsParallel(ctx, subTopics, subTopics[0], &wg)
	}

	for kf.ready.Load() == 0 {
		time.Sleep(10 * time.Millisecond) // Wait and then check again
	}
	util.Log.Info().Msg("Sarama consumer up and running!...")

	sigusr1 := make(chan os.Signal, 1)
	signal.Notify(sigusr1, syscall.SIGUSR1)

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	for keepRunning {
		select {
		case <-sigterm:
			util.Log.Info().Msg("terminating: via signal")
			keepRunning = false
		}
	}

	cancel()
	wg.Wait()
	return nil
}

func (kf *KafkaPubSub) startConsumerGroupsParallel(ctx context.Context, subTopics []string, mainTopic string, wg *sync.WaitGroup) error {
	defer wg.Done()
	consumer, err := sarama.NewConsumerGroupFromClient(fmt.Sprintf("%s-%s", config.Spec.KafkaGroup, mainTopic), kf.kfkClient)
	if err != nil {
		util.Log.Err(err).Msg("error creating consumer group")
		return err
	}

	defer func() {
		if err = consumer.Close(); err != nil {
			util.Log.Panic().Err(err).Msg("error closing client")
		}
	}()

	ctx, cancel := context.WithCancel(ctx)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := consumer.Consume(ctx, subTopics, kf); err != nil {
				if errors.Is(err, sarama.ErrClosedConsumerGroup) {
					return
				}
				util.Log.Err(err).Msg("error from consumer")
			}
			// check if context was cancelled, signaling that the consumer should stop
			if ctx.Err() != nil {
				return
			}
			kf.ready.Store(0)
		}
	}()

	keepRunning := true
	for keepRunning {
		select {
		case <-ctx.Done():
			util.Log.Info().Msg("terminating: via context")
			keepRunning = false
		}
	}

	cancel()
	return nil
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (kf *KafkaPubSub) Setup(sarama.ConsumerGroupSession) error {
	// Mark the consumer as ready
	kf.ready.Store(1)
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (kf *KafkaPubSub) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (kf *KafkaPubSub) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for {
		select {
		case message, ok := <-claim.Messages():
			if !ok {
				util.Log.Info().Msg("claim messages channel closed")
				return nil
			}

			kf.handler.Range(func(key, value any) bool {
				b, err := regexp.MatchString(key.(string), message.Topic)
				if err != nil {
					util.Log.Error().Err(err).Msg("Failed to match topic")
					return true
				}

				if b {
					for _, fn := range value.([]MessageHandler) {
						msg := new(models.Message)
						util.Log.Info().Msg(fmt.Sprintf("Message received: %s", string(message.Value)))
						msg.Topic = message.Topic
						msg.Data = string(message.Value)

						if err := fn(session.Context(), msg); err != nil {
							util.Log.Error().Err(err).Msg("Failed to run handler")
							continue
						}
					}
				}
				return true
			})

			session.MarkMessage(message, fmt.Sprintf("Read on: %s", time.Now().String()))
		// Should return when `session.Context()` is done.
		// If not, will raise `ErrRebalanceInProgress` or `read tcp <ip>:<port>: i/o timeout` when kafka rebalance. see:
		// https://github.com/IBM/sarama/issues/1192
		case <-session.Context().Done():
			return nil
		}
	}
}

package consumer

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/google/uuid"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type SaramaConsumer struct {
	client sarama.ConsumerGroup
	deserializer deserializer
	ready chan bool
}

type deserializer interface {
	Deserialize(value []byte) (string, error)
}

func (c *SaramaConsumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *SaramaConsumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *SaramaConsumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if deserializedMsg, err := c.deserializer.Deserialize(message.Value); err != nil {
			return err
		} else {
			fmt.Println(deserializedMsg)
		}
	}
	return nil
}

func (c *SaramaConsumer) Poll(topic string) error {
	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := c.client.Consume(ctx, []string{topic}, c); err != nil {
				panic(err)
			}

			if ctx.Err() != nil {
				return
			}
			c.ready = make(chan bool)
		}
	}()

	<-c.ready // Await till the consumer has been set up
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
		log.Println("terminating: context cancelled")
	case <-sigterm:
		log.Println("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err := c.client.Close(); err != nil {
		return err
	}

	return nil
}

func NewSaramaConsumer(brokers []string, fromBeginning bool, deserializer deserializer, kafkaVersion string) (*SaramaConsumer, error) {
	cfg := sarama.NewConfig()
	parsedKafkaVersion, err := sarama.ParseKafkaVersion(kafkaVersion)
	if err != nil {
		return nil, err
	}
	cfg.Version = parsedKafkaVersion

	if fromBeginning {
		cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	client, err := sarama.NewConsumerGroup(brokers, "console" + uuid.New().String(), cfg)
	if err != nil {
		return nil, err
	}

	return &SaramaConsumer{
		client:       client,
		deserializer: deserializer,
		ready:        make(chan bool),
	}, nil
}

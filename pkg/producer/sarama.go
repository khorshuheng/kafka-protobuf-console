package producer

import (
	"github.com/Shopify/sarama"
	"github.com/golang/protobuf/proto"
)

type SaramaProducer struct {
	saramaClient sarama.SyncProducer
}

func NewSaramaProducer(brokers []string) (SaramaProducer, error) {
	cfg := sarama.NewConfig()
	cfg.Producer.Return.Successes = true
	client, err := sarama.NewSyncProducer(brokers, cfg)

	return SaramaProducer{client}, err
}

func (p SaramaProducer) Send(topic string, msg proto.Message) error {
	msgByte, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	encodedMsgByte := sarama.ByteEncoder(msgByte)

	pmsg := &sarama.ProducerMessage{
		Topic:	topic,
		Value:	encodedMsgByte,
	}

	_, _, err = p.saramaClient.SendMessage(pmsg)
	return err
}

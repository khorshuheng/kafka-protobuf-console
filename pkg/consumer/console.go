package consumer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/khorshuheng/kafka-protobuf-console/pkg/configs"
	"github.com/khorshuheng/kafka-protobuf-console/pkg/reflection"
)


type kafkaConsumer interface {
	Poll(topic string) error
}

type Console struct {
	kafkaConsumer kafkaConsumer
	topic string
}

type ProtoDeserializer struct {
	msgDescriptor *desc.MessageDescriptor
}

func (pd *ProtoDeserializer) Deserialize(value []byte) (string, error) {
	dymsg := dynamic.NewMessage(pd.msgDescriptor)
	proto.Unmarshal(value, dymsg)
	marshaller := jsonpb.Marshaler{}
	if jsonStr, err := marshaller.MarshalToString(dymsg); err != nil {
		return "", err
	} else {
		return jsonStr, nil
	}
}

func NewConsole(cfg config.Consumer) (*Console, error) {
	md, err := reflection.LoadMessageDescriptor(cfg.FileDescriptorPath, cfg.ProtoName)
	if err != nil {
		return nil, err
	}

	kc, err := NewSaramaConsumer(cfg.Brokers, cfg.FromBeginning, &ProtoDeserializer{md})
	if err != nil {
		return nil, err
	}

	return &Console{
		kafkaConsumer: kc,
		topic: cfg.Topic,
	}, nil
}

func (c *Console) Start() {
	c.kafkaConsumer.Poll(c.topic)
}

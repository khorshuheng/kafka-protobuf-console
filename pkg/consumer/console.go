package consumer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/khorshuheng/kafka-protobuf-console/pkg/config"
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
	prettyPrint bool
}

func (pd *ProtoDeserializer) Deserialize(value []byte) (string, error) {
	dymsg := dynamic.NewMessage(pd.msgDescriptor)
	err := proto.Unmarshal(value, dymsg)
	if err != nil {
		return "", err
	}
	indent := ""
	if pd.prettyPrint {
		indent = "  "
	}
	marshaller := jsonpb.Marshaler{Indent: indent}
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

	protoDeserializer := &ProtoDeserializer{md, cfg.PrettyPrint}
	kc, err := NewSaramaConsumer(cfg.Brokers, cfg.FromBeginning, protoDeserializer, cfg.Version)
	if err != nil {
		return nil, err
	}

	return &Console{
		kafkaConsumer: kc,
		topic: cfg.Topic,
	}, nil
}

func (c *Console) Start() error {
	err := c.kafkaConsumer.Poll(c.topic)
	if err != nil {
		return err
	}
	return nil
}

package producer

import (
	"bufio"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
	"github.com/khorshuheng/kafka-protobuf-console/pkg/config"
	"github.com/khorshuheng/kafka-protobuf-console/pkg/reflection"
	"os"
)

type kafkaProducer interface {
	Send(topic string, msg proto.Message) error
}

type Console struct {
	kafkaProducer kafkaProducer
	msgDescriptor *desc.MessageDescriptor
	topic		  string
}

func NewConsole(cfg config.Producer) (*Console, error) {
	kp, err := NewSaramaProducer(cfg.Brokers)
	if err != nil {
		return nil, err
	}
	md, err := reflection.LoadMessageDescriptor(cfg.FileDescriptorPath, cfg.ProtoName)
	if err != nil {
		return nil, err
	}
	return &Console{
		kafkaProducer: kp,
		msgDescriptor: md,
		topic:         cfg.Topic,
	}, nil
}

func (c *Console) Start() error {
	scanner := bufio.NewScanner(os.Stdin)
	promptInput()
	for scanner.Scan() {
		err := c.processInput(scanner.Text())
		if err != nil {
			return err
		}
		promptInput()
	}

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}

func (c *Console) processInput(userInput string) error {
	dymsg := dynamic.NewMessage(c.msgDescriptor)
	if err := jsonpb.UnmarshalString(userInput, dymsg); err != nil {
		return err
	}
	return c.kafkaProducer.Send(c.topic, dymsg)
}

func promptInput() {
	fmt.Print("> ")
}

package producer

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockKafkaProducer struct {
	mock.Mock
}

func (p MockKafkaProducer) Send(topic string, msg proto.Message) error {
	args := p.Called(topic, msg)
	return args.Error(0)
}

func TestConsole(suite *testing.T) {
	strType := builder.FieldTypeScalar(descriptor.FieldDescriptorProto_TYPE_STRING)
	msg, _ := builder.NewMessage("MyMessage").AddField(
		builder.NewField("foo", strType),
		).Build()
	topic := "sometopic"
	kp := MockKafkaProducer{}
	kp.On("Send", topic, mock.Anything).Return(nil)
	console := Console{
		kafkaProducer: kp,
		msgDescriptor: msg,
		topic:         topic,
	}

	suite.Run("processInput_ShouldConvertJsonStringToProtobuf", func(t *testing.T) {
		err := console.processInput("{\"foo\": \"bar\"}")
		assert.NoError(t, err)
	})

	suite.Run("processInput_ShouldReturnErrorIfSchemaIsIncorrect", func(t *testing.T) {
		err := console.processInput("{\"bar\": \"foo\"}")
		assert.Error(t, err)
	})

	suite.Run("processInput_ShouldReturnErrorIfInputIsNotJson", func(t *testing.T) {
		err := console.processInput("")
		assert.Error(t, err)
	})
}
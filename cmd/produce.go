package cmd

import (
	"github.com/khorshuheng/kafka-protobuf-console/configs"
	"github.com/khorshuheng/kafka-protobuf-console/pkg/producer"
	"github.com/spf13/cobra"
)

var produceCmd = &cobra.Command{
	Use:   	"produce",
	Short: 	"Produce protobuf message to Kafka using Json input",
	Run:  	produce,
}

func init() {
	produceCmd.Flags().StringP("descriptor", "d", "", "File descriptor path")
	produceCmd.Flags().StringP("name", "n", "", "Fully qualified Proto message name")
	produceCmd.Flags().StringP("topic", "t", "", "Destination Kafka topic")
	produceCmd.MarkFlagRequired("name")
	produceCmd.MarkPersistentFlagRequired("brokers")
}

func produce(cmd *cobra.Command, args []string) {
	brokers, err := cmd.Flags().GetStringSlice("brokers")
	if err != nil {
		panic(err)
	}

	fileDescriptorPath, err := cmd.Flags().GetString("descriptor")
	if err != nil {
		panic(err)
	}

	protoMessageName, err := cmd.Flags().GetString("name")
	if err != nil {
		panic(err)
	}

	console, err := producer.NewConsole(configs.ProducerConfig{
		Brokers:            brokers,
		FileDescriptorPath: fileDescriptorPath,
		ProtoName:          protoMessageName,
	})
	if err != nil {
		panic(err)
	}

	err = console.Start()
	if err != nil {
		panic(err)
	}
}
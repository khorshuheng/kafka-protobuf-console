package cmd

import (
	"github.com/khorshuheng/kafka-protobuf-console/pkg/config"
	"github.com/khorshuheng/kafka-protobuf-console/pkg/producer"
	"github.com/spf13/cobra"
)

var produceCmd = &cobra.Command{
	Use:   	"produce",
	Short: 	"Produce protobuf message to Kafka using Json input",
	Run:  	produce,
}

func init() {
	produceCmd.Flags().StringSliceP("brokers", "b", []string{},
		"Comma separated Kafka brokers address")
	produceCmd.Flags().StringP("descriptor", "d", "", "File descriptor path")
	produceCmd.Flags().StringP("name", "n", "", "Fully qualified Proto message name")
	produceCmd.Flags().StringP("topic", "t", "", "Destination Kafka topic")
	produceCmd.MarkFlagRequired("name")
	produceCmd.MarkFlagRequired("brokers")
	produceCmd.MarkFlagRequired("topic")
}

func produce(cmd *cobra.Command, args []string) {
	commonCfg, err := ParseCommonConfig(cmd)
	if err != nil {
		panic(err)
	}
	console, err := producer.NewConsole(config.Producer{
		Common: commonCfg,
	})
	if err != nil {
		panic(err)
	}

	err = console.Start()
	if err != nil {
		panic(err)
	}
}
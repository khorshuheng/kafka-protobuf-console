package cmd

import (
	"github.com/khorshuheng/kafka-protobuf-console/pkg/configs"
	"github.com/khorshuheng/kafka-protobuf-console/pkg/producer"
	"github.com/spf13/cobra"
)

var produceCmd = &cobra.Command{
	Use:   	"produce",
	Short: 	"Produce protobuf message to Kafka using Json input",
	Run:  	produce,
}

func init() {
	produceCmd.MarkFlagRequired("name")
	produceCmd.MarkPersistentFlagRequired("brokers")
	produceCmd.MarkPersistentFlagRequired("topic")
}

func produce(cmd *cobra.Command, args []string) {
	commonCfg, err := ParseCommonConfig(cmd)
	if err != nil {
		panic(err)
	}
	console, err := producer.NewConsole(configs.Producer{
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
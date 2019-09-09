package cmd

import (
	"github.com/khorshuheng/kafka-protobuf-console/pkg/config"
	"github.com/khorshuheng/kafka-protobuf-console/pkg/consumer"
	"github.com/spf13/cobra"
)

var consumeCmd = &cobra.Command{
	Use:   	"consume",
	Short: 	"Consume Protobuf message from Kafka",
	Run:  	consume,
}

func init() {
	consumeCmd.Flags().BoolP("from-beginning", "f", false, "Consume from beginning")
	consumeCmd.Flags().StringSliceP("brokers", "b", []string{},
		"Comma separated Kafka brokers address")
	consumeCmd.Flags().StringP("descriptor", "d", "", "File descriptor path")
	consumeCmd.Flags().StringP("name", "n", "", "Fully qualified Proto message name")
	consumeCmd.Flags().StringP("topic", "t", "", "Destination Kafka topic")
	consumeCmd.Flags().StringP("version", "v", "", "Kafka version (eg. 2.0.0)")
	consumeCmd.MarkFlagRequired("name")
	consumeCmd.MarkFlagRequired("brokers")
	consumeCmd.MarkFlagRequired("topic")
	consumeCmd.MarkFlagRequired("version")
}

func consume(cmd *cobra.Command, args []string) {
	commonCfg, err := ParseCommonConfig(cmd)
	if err != nil {
		panic(err)
	}

	fromBeginning, err := cmd.Flags().GetBool("from-beginning")
	if err != nil {
		panic(err)
	}

	kafkaVersion, err := cmd.Flags().GetString("version")
	if err != nil {
		panic(err)
	}

	consumerCfg := config.Consumer{
		Common:        commonCfg,
		FromBeginning: fromBeginning,
		Version:	   kafkaVersion,
	}

	console, err := consumer.NewConsole(consumerCfg)
	if err != nil {
		panic(err)
	}
	console.Start()
}

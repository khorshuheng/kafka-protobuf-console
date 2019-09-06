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
	consumeCmd.MarkFlagRequired("name")
	consumeCmd.MarkPersistentFlagRequired("brokers")
	consumeCmd.MarkPersistentFlagRequired("topic")
}

func consume(cmd *cobra.Command, args []string) {
	commonCfg, err := ParseCommonConfig(cmd)
	if err != nil {
		panic(err)
	}

	fromBeginning, err := cmd.Flags().GetBool("from-beginning")
	consumerCfg := config.Consumer{
		Common:        commonCfg,
		FromBeginning: fromBeginning,
	}

	console, err := consumer.NewConsole(consumerCfg)
	if err != nil {
		panic(err)
	}
	console.Start()
}

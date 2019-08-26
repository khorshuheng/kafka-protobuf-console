package cmd

import (
	"github.com/khorshuheng/kafka-protobuf-console/configs"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kafka-protobuf-console",
	Short: "Utility to produce / consume console message to / from Kafka",
}

func init() {
	cobra.OnInitialize()
	rootCmd.PersistentFlags().StringSliceP("brokers", "b", []string{},
	"Comma separated Kafka brokers address")
	rootCmd.PersistentFlags().StringP("descriptor", "d", "", "File descriptor path")
	rootCmd.PersistentFlags().StringP("name", "n", "", "Fully qualified Proto message name")
	rootCmd.PersistentFlags().StringP("topic", "t", "", "Destination Kafka topic")
	rootCmd.AddCommand(produceCmd)
	rootCmd.AddCommand(consumeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

func ParseCommonConfig(cmd *cobra.Command) (configs.Common, error) {
	brokers, err := cmd.Flags().GetStringSlice("brokers")
	if err != nil {
		return configs.Common{}, err
	}

	fileDescriptorPath, err := cmd.Flags().GetString("descriptor")
	if err != nil {
		return configs.Common{}, err
	}

	protoMessageName, err := cmd.Flags().GetString("name")
	if err != nil {
		return configs.Common{}, err
	}

	topic, err := cmd.Flags().GetString("topic")
	if err != nil {
		return configs.Common{}, err
	}

	return configs.Common{
		Brokers:            brokers,
		FileDescriptorPath: fileDescriptorPath,
		ProtoName:          protoMessageName,
		Topic:              topic,
	}, nil
}
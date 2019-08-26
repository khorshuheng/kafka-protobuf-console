package cmd

import (
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
	rootCmd.MarkPersistentFlagRequired("brokers")
	rootCmd.AddCommand(produceCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
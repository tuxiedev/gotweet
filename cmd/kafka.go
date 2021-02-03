/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tuxiedev/gotweet/pkg/structs"
)

// kafkaCmd represents the kafka command
var kafkaCmd = &cobra.Command{
	Use:   "kafka",
	Short: "Produce tweets to Kafka",
	Long:  `When running in this mode, the application produces incoming tweets to a Kafka topic`,
	Run: func(cmd *cobra.Command, args []string) {
		runApp("kafka", getKafkaConfiguration())
	},
}

var outputTopic string
var bootstrapBrokers []string

func getKafkaConfiguration() structs.KafkaConfig {
	fmt.Println("cmd...", bootstrapBrokers)
	return structs.KafkaConfig{
		BootstrapBrokers: bootstrapBrokers,
		OutputTopic:      outputTopic,
	}
}

func init() {
	rootCmd.AddCommand(kafkaCmd)

	pf := kafkaCmd.Flags()

	pf.StringVarP(&outputTopic, "output-topic", "o", viper.GetString("KAFKA_OUTPUT_TOPIC"), "Kafka output topic")
	pf.StringArrayVarP(&bootstrapBrokers, "bootstrap-brokers", "b", viper.GetStringSlice("KAFKA_BOOTSTRAP_BROKERS"), "REQUIRED: Comma separated list of kafka brokers to connect to")

	cobra.MarkFlagRequired(pf, "bootstrap-brokers")

}

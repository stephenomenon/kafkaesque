package main

import (
	"fmt"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

const (
	version = "version"
	brokers = "brokers"
	topics  = "topics"
	group   = "group"
	clientid = "clientid"
)

type ConfigurationError string

func (err ConfigurationError) Error() string {
	return "configuration error (" + string(err) + ")"
}

func ConfigFromFile(updateCh chan sig) error {

	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		updateCh<-sig{}
	})

	return nil
}

func verifyConfig() error {
	switch {
	case ConsumerGroup() == "":
		return ConfigurationError("missing ConsumerGroup")

	case ClientID() == "":
		return ConfigurationError("missing ClientID")

	case len(kafkaBrokers()) == 0:
		return ConfigurationError("missing Kafka Brokers")

	case len(ConsumerTopics()) == 0:
		return ConfigurationError("missing Consumer Topics")

	}

	return nil
}

func displayConfig() {

	fmt.Printf("\nusing sarama consumer version: %s\n", viper.GetString("version"))

	g := ConsumerGroup()
	fmt.Printf("using consumer group: %s\n", g)

	id := ClientID()
	fmt.Printf("using client id: %s\n", id)

	b := kafkaBrokers()
	fmt.Printf("using kafka brokers: %s\n", strings.Join(b[:], ","))

	t := ConsumerTopics()
	fmt.Printf("using kafka topics: %s\n\n", strings.Join(t[:], ","))
}

func kafkaVersion() sarama.KafkaVersion {
	v := viper.GetString(version)

	switch v {
	case "0.10.2.0":
		return sarama.V0_10_2_0

	case "0.10.1.0":
		return sarama.V0_10_1_0

	case "0.10.0.1":
		return sarama.V0_10_0_1

	case "0.10.0.0":
		return sarama.V0_10_0_0

	case "0.9.0.1":
		return sarama.V0_9_0_1

	case "0.9.0.0":
		return sarama.V0_9_0_0

	default:
		return sarama.V0_8_2_2
	}
}

func ConsumerGroup() string {
	return viper.GetString(group)
}

func ClientID() string {
	return viper.GetString(clientid)
}

func kafkaBrokers() []string {
	return viper.GetStringSlice(brokers)
}

func ConsumerTopics() []string {
	return viper.GetStringSlice(topics)
}

package main

import (
	"github.com/spf13/viper"
	"fmt"
	"github.com/Shopify/sarama"
	"errors"
	"strings"
)

const (
	version = "version"
	brokers = "brokers"
	topics = "topics"
)

func KafkaConfigFromEnv() {
	viper.SetEnvPrefix("kafka") // will be uppercased automatically
	viper.BindEnv(version)
	viper.BindEnv(brokers)
	viper.BindEnv(topics)

}

func KafkaConfigFromFile() {

	viper.SetConfigType("toml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

func displayConfig() {

	fmt.Printf("using sarama consumer version: %s\n", viper.GetString("version"))

	b := kafkaBrokers()
	fmt.Printf("using kafka brokers: %s\n", strings.Join(b[:],","))

	t := kafkaTopics()
	fmt.Printf("using kafka topics: %s\n", strings.Join(t[:],","))
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

func kafkaBrokers() []string {
	b := viper.GetStringSlice(brokers)
	if len(b) == 0 {
		err := errors.New("No kafka brokers set in config")
		panic(fmt.Errorf("\nFatal error config: %s \n", err))
	}

	return b
}

func kafkaTopics() []string{
	t := viper.GetStringSlice(topics)
	if len(t) == 0 {
		err := errors.New("No kafka topics set in config")
		panic(fmt.Errorf("\nFatal error config: %s \n", err))
	}

	return t
}
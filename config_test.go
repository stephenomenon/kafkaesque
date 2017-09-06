package main

import (
	"testing"
	"github.com/spf13/viper"
	"os"
)

func TestConfigFromEnv(t *testing.T) {

	os.Setenv("KAFKA_VERSION", "0.10.2.0")
	os.Setenv("KAFKA_BROKERS", "['10.0.0.1','10.0.0.2','10.0.0.3']")
	os.Setenv("KAFKA_TOPICS", "['test-topic-1','test-topic-2']")

	KafkaConfigFromEnv()
	displayConfig()

}

func TestConfigFromFile(t *testing.T) {

	ConfigFromFile()

	t.Log(viper.ConfigFileUsed())
	t.Log(viper.AllSettings())

}
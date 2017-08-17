package main

import (
	"testing"
)

func TestKafkaConsumer(t *testing.T) {

	KafkaConfigFromEnv()
	displayConfig()

	startKafkaConsumer()

	// TODO: get some assertions going on here
	// TODO: try the sarama mock package
}
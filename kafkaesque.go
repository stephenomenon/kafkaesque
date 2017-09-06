package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
	cluster "github.com/bsm/sarama-cluster"

	"time"
)

// startKafkaConsumer is taken directly from the sarama-cluster example
func KafkaConsumer(stopCh, doneCH chan sig) {

	// init (custom) config, enable errors and notifications
	config := cluster.NewConfig()
	config.Version = kafkaVersion()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	config.ClientID = ClientID()

	// init consumer
	brokers := kafkaBrokers()
	group := ConsumerGroup()
	topics := ConsumerTopics()

	consumer, err := cluster.NewConsumer(brokers, group, topics, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		consumer.Close()
		doneCH <- sig{}
	}()

	// consume messages, watch errors and notifications
	for {
		select {
		case msg, more := <-consumer.Messages():
			if more {
				fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				// simulate some work
				time.Sleep(200*time.Millisecond)
				consumer.MarkOffset(msg, "") // mark message as processed
			}
		case err, more := <-consumer.Errors():
			if more {
				log.Printf("Error: %s\n", err.Error())
			}
		case ntf, more := <-consumer.Notifications():
			if more {
				log.Printf("Rebalanced: %+v\n", ntf)
			}

		case <-stopCh:
			log.Printf("\t...stoping consumer group")
			return

		}
	}
}

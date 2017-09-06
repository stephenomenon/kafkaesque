package main

import (
	"os"
	"os/signal"
	"fmt"
)

type sig struct{}

func main() {

	updateCh := make(chan sig)
	if err := ConfigFromFile(updateCh); err != nil {
		panic(err)
	}

	os.Exit(runLoop(updateCh))
}

func runLoop(ConfigCh <-chan sig) int {

	// trap SIGINT to trigger a shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
	defer close(signals)

	for {

		// verify config
		if err := verifyConfig(); err != nil {
			fmt.Println(err.Error())
			return -1
		}

		displayConfig()

		stopCh, doneCh := make(chan sig), make(chan sig)
		go KafkaConsumer(stopCh, doneCh)

		select {
		case <-ConfigCh:
			fmt.Println("Config file changed!")
			close(stopCh)
			<-doneCh
			// restart the run loop

		case <-signals:
			close(stopCh)
			<-doneCh
			return 0
		}
	}

	return -1
}


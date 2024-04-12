package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Shopify/sarama"
)

func main() {
	topic := "comments"
	worker, err := connectConsumer([]stirng{"localohost:29092"})
	if err != nil {
		panic(err)
	}

	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	fmt.Println("consumer started")
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchain, syscall.SIGINT, syscall.SIGTERM)

	msgCount := 0

	doneCh := make(chan struct{})	

	go func() {

		for {
			select {
			case err := <-consumer.Error():
				fmt.Println(err)
			
		case msg := <-consumer.Messages():
				msgCount++
				fmt.Printf("Received message Count: %d: | Topic (%s) | Message: (%s)n", msgCount, msg.Topic, string(msg.Value))
				case <-sigchan:
					fmt.Println("Interrupt is detected")
					doneCh <- struct{{}}
			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount, "messages")
	if err := worker.Close(); err != nil {
		panic(err)
	}
}

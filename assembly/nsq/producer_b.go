package main

import (
	"fmt"
	"log"
	"time"

	"github.com/nsqio/go-nsq"
)

func main() {
	// Instantiate a producer.
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer("127.0.0.1:4150", config)
	if err != nil {
		log.Fatal(err)
	}

	topicName := "my_topic_test"

	// Synchronously publish a single message to the specified topic.
	// Messages can also be sent asynchronously and/or in batches.
	var process = make(chan int)
	for i := 1; i <= 5; i++ {
		go func(process chan int) {
			for i := 0; i < 99999; i++ {
				messageBody := []byte(fmt.Sprintf("hello %d,time:%s", i, time.Now().Format("2006-01-02 15:04:05")))
				err = producer.Publish(topicName, messageBody)
				if err != nil {
					log.Fatal(err)
				}
				time.Sleep(time.Second * time.Duration(1))
			}
			process <- 1
		}(process)
		<-process
	}

	// Gracefully stop the producer when appropriate (e.g. before shutting down the service)
	producer.Stop()
}

package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"html"
	"log"
	"net/http"
	"time"
)

var brokers = []string{"localhost:9092"}

const topic = "demo-topic"

var fakeDB string

func saveMessage(text string) { fakeDB = text }

func getMessage() string { return fakeDB }

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Println("producer err: ", err)
	}

	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		fmt.Println("Could not create consumer: ", err)
	}

	subscribe(topic, consumer)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, "Hello Sarama!") })

	http.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		r.ParseForm()
		msg := prepareMessage(topic, r.FormValue("q"))
		partition, offset, err := producer.SendMessage(msg)
		fmt.Fprintf(w, "Message was saved to partition: %d.\nMessage offset is: %d.\n.", partition, offset)
		if err != nil {
			fmt.Fprintf(w, "Error occurred - %s.\n", err)
		}
	})
	http.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, html.EscapeString(getMessage())) })

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	fmt.Println(message)
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
		Timestamp: time.Now(),
	}
	return msg
}

func subscribe(topic string, consumer sarama.Consumer) {
	partitionList, err := consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		fmt.Println("Error retrieving partitionList ", err)
	}
	initialOffset := sarama.OffsetOldest //get offset for the oldest message on the topic

	for _, partition := range partitionList {
		pc, _ := consumer.ConsumePartition(topic, partition, initialOffset)

		go func(pc sarama.PartitionConsumer) {
			for message := range pc.Messages() {
				messageReceived(message)
			}
		}(pc)
	}
}

func messageReceived(message *sarama.ConsumerMessage) {
	saveMessage(string(message.Value))
}

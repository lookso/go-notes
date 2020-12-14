package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"html"
	"log"
	"net/http"
	"strconv"
	"sync"
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
	config.Producer.Retry.Max = 3
	config.Producer.Timeout = 5 * time.Second
	config.Producer.Return.Successes = true
	config.Version = sarama.V0_10_2_1 // 指定kafka 版本,不然写入timestamp字段为空
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		fmt.Println("producer err: ", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	// 消费者
	go func(brokers []string) {
		defer wg.Done()
		consumer, err := sarama.NewConsumer(brokers, nil)
		config := sarama.NewConfig()
		if err != nil {
			fmt.Println("Could not create consumer: ", err)
		}
		subscribe(topic, consumer)
	}(brokers)
	wg.Wait()

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
	dt := map[string]interface{}{
		"message":   message,
		"timestamp": time.Now(),
		"meg":       "20201214-test-go-kafka",
	}
	value, _ := json.Marshal(&dt)
	msg := &sarama.ProducerMessage{
		Key:       sarama.StringEncoder(strconv.FormatInt(time.Now().UTC().UnixNano(), 10)),
		Topic:     topic,
		Partition: -1,
		Value:     sarama.ByteEncoder(value),
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
	fmt.Println("consumer-received-value", string(message.Value))
	saveMessage(string(message.Value))
}

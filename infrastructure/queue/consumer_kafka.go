package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go-dummy-clean-arch/api/presenter"
	"go-dummy-clean-arch/usecase/artist"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type kafkaConsumer struct {
	topic    string
	sigchan  chan os.Signal
	consumer *kafka.Consumer
	service  artist.UseCase
}

// https://developer.confluent.io/get-started/go/#build-consumer

func NewKafkaConsumer(topic string, service artist.UseCase) Consumer {
	conf := kafka.ConfigMap{
		"bootstrap.servers": os.Getenv("KAFKA_BOOTSTRAP"),
		"group.id":          "kafka-go-getting-started",
		"auto.offset.reset": "earliest",
	}

	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		// TODO: return error
		fmt.Printf("Failed to create consumer: %s", err)
		os.Exit(1)
	}

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		// TODO: return error
	}

	return &kafkaConsumer{
		topic:    topic,
		sigchan:  make(chan os.Signal, 1),
		consumer: consumer,
		service:  service,
	}
}

func (k *kafkaConsumer) Consume() error {
	// Set up a channel for handling Ctrl-C, etc
	signal.Notify(k.sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	run := true
	for run == true {
		select {
		case sig := <-k.sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := k.consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}
			k.save(string(ev.Key), string(ev.Value))
			fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
				*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))
		}
	}

	k.consumer.Close()
	return nil
}

func (k *kafkaConsumer) save(key, value string) {
	artistJson := &presenter.Artist{}
	err := json.Unmarshal([]byte(value), artistJson)
	if err != nil {
		fmt.Println(err)
		return
	}

	artist := artistJson.ToEntity()
	err = k.service.Save(context.Background(), *artist)
	// TODO: handle error
}

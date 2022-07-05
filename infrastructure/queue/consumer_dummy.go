package queue

import (
	"log"
)

type dummyConsumer struct {
	topic string
}

func NewDummyConsumer(topic string) Consumer {
	return &dummyConsumer{topic: topic}
}

func (d *dummyConsumer) Consume() error {
	log.Println("Dummy Consumer consuming topic:", d.topic)
	return nil
}

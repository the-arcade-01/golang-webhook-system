package main

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	addr   string
	topic  string
	writer *kafka.Writer
}

func NewKafkaClient(addr, topic string) *KafkaClient {
	return &KafkaClient{
		addr:  addr,
		topic: topic,
		writer: &kafka.Writer{
			Addr:                   kafka.TCP(addr),
			Topic:                  topic,
			AllowAutoTopicCreation: true,
		},
	}
}

func (c *KafkaClient) PublishEvent(ctx context.Context, event Event) error {
	eventJSON, err := json.Marshal(event)
	if err != nil {
		return err
	}

	return c.writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(event.EventID),
		Value: eventJSON,
	})
}

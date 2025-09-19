package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/segmentio/kafka-go"
)

type KafkaClient struct {
	addr       string
	topic      string
	groupId    string
	reader     *kafka.Reader
	writer     *kafka.Writer
	db         *EventsDB
	maxRetries int
}

func NewKafkaClient(addr, topic, groupId string, db *EventsDB) *KafkaClient {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        []string{addr},
		Topic:          topic,
		GroupID:        groupId,
		CommitInterval: 0,
	})

	writer := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}

	return &KafkaClient{
		addr:       addr,
		topic:      topic,
		groupId:    groupId,
		reader:     reader,
		writer:     writer,
		db:         db,
		maxRetries: 3,
	}
}

func (c *KafkaClient) ConsumeEvents(ctx context.Context) {
	log.Println("started consuming events from kafka")

	for {
		m, err := c.reader.FetchMessage(ctx)
		if err != nil {
			log.Printf("error reading from kafka topic. %s", err)
			break
		}

		var msg Event
		if err := json.Unmarshal(m.Value, &msg); err != nil {
			log.Printf("error parsing event from kafka. %s\n", err)
			continue
		}

		if rand.Intn(10) < 3 {
			log.Printf("simulating event processing failure for eventID: %s", msg.EventID)
			c.handleFailure(ctx, m)
			continue
		}

		err = c.db.Insert(ctx, msg)
		if err != nil {
			log.Printf("error inserting into db. %s", err)
			continue
		}

		if err := c.reader.CommitMessages(ctx, m); err != nil {
			log.Printf("error committing msg offset, %s", err)
		} else {
			log.Printf("event successfully inserted into db. eventID: %s\n", msg.EventID)
		}
	}

	if err := c.reader.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}

func (c *KafkaClient) handleFailure(ctx context.Context, msg kafka.Message) {
	retryCount := 0
	for _, h := range msg.Headers {
		if h.Key == "retry_count" {
			val, _ := strconv.Atoi(string(h.Value))
			retryCount = val
		}
	}
	retryCount++

	if retryCount > c.maxRetries {
		log.Printf("moving to DLQ: %s", string(msg.Value))
		if err := c.reader.CommitMessages(ctx, msg); err != nil {
			log.Printf("commit after DLQ failed: %v", err)
		}
		return
	}

	// Push to retry topic with incremented count
	log.Printf("retrying event: %s, attempt %d", string(msg.Value), retryCount)
	newHeaders := append(msg.Headers, kafka.Header{Key: "retry_count", Value: fmt.Appendf(nil, "%d", retryCount)})

	c.writer.WriteMessages(ctx, kafka.Message{
		Value:   msg.Value,
		Headers: newHeaders,
	})
}

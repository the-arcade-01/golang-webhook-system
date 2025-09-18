package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
)

type HTTPClientInterface interface {
	Do(*http.Request) (*http.Response, error)
}

// Event represents a webhook event.
type Event struct {
	EventID   string            `json:"event_id"`
	EventType WebhookEventTypes `json:"event_type"`
	Timestamp time.Time         `json:"timestamp"`
	Data      EventData         `json:"data"`
}

// EventData holds additional event data.
type EventData struct {
	InvoiceID string `json:"invoice_id"`
}

type Job struct {
	Ticker         *time.Ticker
	Quit           chan struct{}
	Key            string
	WebhookDetails WebhookDetails
}

// Emitter is a simple event emitter class to simulate event publishing in a webhook system.
// It periodically generates random events and sends them to configured webhook URLs.
type Emitter struct {
	jobs     map[string]*Job
	db       *WebhookDB
	interval int
	mtx      sync.Mutex
	client   HTTPClientInterface
}

func NewEmitter(db *WebhookDB, client HTTPClientInterface, cron int) *Emitter {
	return &Emitter{
		db:       db,
		interval: cron,
		client:   client,
	}
}

func (e *Emitter) Start(customerID, webhookID string) error {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	webhookDetails, err := e.db.Get(context.Background(), customerID, webhookID)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(time.Duration(e.interval) * time.Second)
	quitCh := make(chan struct{})
	key := e.getKey(customerID, webhookID)

	job := &Job{
		Ticker:         ticker,
		Quit:           quitCh,
		Key:            key,
		WebhookDetails: webhookDetails,
	}
	e.jobs[job.Key] = job

	go func() {
		for {
			select {
			case <-job.Ticker.C:
				event := e.randomEventGenerator()
				eventJSON, err := json.Marshal(event)
				if err != nil {
					log.Printf("error on marshaling event, %s", err)
					continue
				}

				req, err := http.NewRequest("POST", job.WebhookDetails.WebhookURL, bytes.NewBuffer(eventJSON))
				if err != nil {
					log.Printf("error creating request: %s", err)
					continue
				}
				req.Header.Set("Content-Type", "application/json")
				req.Header.Set("x-webhook-signature", GenerateHMAC(eventJSON, job.WebhookDetails.SecretToken))

				resp, err := e.client.Do(req)
				if err != nil {
					log.Printf("error sending webhook: %s", err)
					continue
				}
				resp.Body.Close()
				log.Printf("event sent successfully, event_id: %s\n", event.EventID)

			case <-job.Quit:
				job.Ticker.Stop()
				log.Printf("Event emitter stopped for %s", job.Key)
				return
			}
		}
	}()

	return nil
}

func (e *Emitter) Stop(customerID, webhookID string) {
	e.mtx.Lock()
	defer e.mtx.Unlock()

	key := e.getKey(customerID, webhookID)
	if job, ok := e.jobs[key]; ok {
		close(job.Quit)
		delete(e.jobs, key)
	}
}

func (e *Emitter) getKey(customerID, webhookID string) string {
	return fmt.Sprintf("%s-%s", customerID, webhookID)
}

func (e *Emitter) randomEventGenerator() Event {
	eventTypes := []WebhookEventTypes{CREATED, PENDING, PAID, FAILED}

	return Event{
		EventID:   uuid.New().String(),
		EventType: eventTypes[rand.Intn(len(eventTypes))],
		Timestamp: time.Now(),
		Data: EventData{
			InvoiceID: uuid.New().String(),
		},
	}
}

package main

import (
	"errors"
	"time"
)

/** For simplicity we are keeping predefined event types
 */
type WebhookEventTypes int

const (
	CREATED WebhookEventTypes = iota
	PENDING
	PAID
	FAILED
)

func (t WebhookEventTypes) ToString() string {
	switch t {
	case CREATED:
		return "CREATED"
	case PENDING:
		return "PENDING"
	case PAID:
		return "PAID"
	case FAILED:
		return "FAILED"
	default:
		return "INVALID"
	}
}

func (t WebhookEventTypes) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.ToString() + `"`), nil
}

func (t *WebhookEventTypes) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"CREATED"`:
		*t = CREATED
	case `"PENDING"`:
		*t = PENDING
	case `"PAID"`:
		*t = PAID
	case `"FAILED"`:
		*t = FAILED
	default:
		return errors.New("invalid enum type for WebhookEventTypes")
	}
	return nil
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

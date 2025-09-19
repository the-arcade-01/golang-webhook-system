package main

import (
	"errors"
	"fmt"
	"time"
)

/** WebhookStatus
 */
type WebhookStatus int

const (
	ACTIVE WebhookStatus = iota
	DISABLED
)

func (t WebhookStatus) ToString() string {
	switch t {
	case ACTIVE:
		return "ACTIVE"
	case DISABLED:
		return "DISABLED"
	default:
		return "INVALID"
	}
}

func (t WebhookStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.ToString() + `"`), nil
}

func (t *WebhookStatus) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"ACTIVE"`:
		*t = ACTIVE
	case `"DISABLED"`:
		*t = DISABLED
	default:
		return errors.New("invalid enum type for WebhookStatus")
	}
	return nil
}

func (t *WebhookStatus) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid data type for WebhookStatus")
	}

	switch str {
	case "ACTIVE":
		*t = ACTIVE
	case "DISABLED":
		*t = DISABLED
	default:
		return fmt.Errorf("invalid enum type for WebhookStatus: %s", str)
	}
	return nil
}

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

type WebhookRegisterBody struct {
	CustomerID string `json:"customer_id"`
	WebhookURL string `json:"webhook_url"`
}

type WebhookDetails struct {
	WebhookID     string        `json:"webhook_id"`
	WebhookURL    string        `json:"webhook_url"`
	WebhookStatus WebhookStatus `json:"webhook_status"`
	SecretToken   string        `json:"secret_token"`
}

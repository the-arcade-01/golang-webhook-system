package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	client *KafkaClient
}

func NewHandler(client *KafkaClient) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) RequestHandler(w http.ResponseWriter, r *http.Request) {
	var body Event
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid event object", http.StatusBadRequest)
		return
	}
	err := h.client.PublishEvent(r.Context(), body)
	if err != nil {
		log.Printf("error sending event to kafka, %s", err)
		http.Error(w, "Event processing failed", http.StatusInternalServerError)
		return
	}

	log.Printf("event object sent to kafka, %v", body.EventID)
	w.WriteHeader(http.StatusOK)
}

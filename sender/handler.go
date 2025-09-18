package main

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	svc     *Service
	emitter *Emitter
}

func NewHandler(svc *Service, emitter *Emitter) *Handler {
	return &Handler{
		svc:     svc,
		emitter: emitter,
	}
}

func (h *Handler) RegisterWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	var body WebhookRegisterBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid body passed", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	res, err := h.svc.RegisterWebhook(r.Context(), body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) GetWebhooks(w http.ResponseWriter, r *http.Request) {
	res, err := h.svc.db.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) StartEmitter(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	customerID := r.URL.Query().Get("customer_id")
	webhookID := r.URL.Query().Get("webhook_id")

	if customerID == "" || webhookID == "" {
		http.Error(w, "Invalid params passed", http.StatusBadRequest)
		return
	}

	err := h.emitter.Start(customerID, webhookID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Event emitter started")
}

func (h *Handler) StopEmitter(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid method", http.StatusBadRequest)
		return
	}

	customerID := r.URL.Query().Get("customer_id")
	webhookID := r.URL.Query().Get("webhook_id")

	if customerID == "" || webhookID == "" {
		http.Error(w, "Invalid params passed", http.StatusBadRequest)
		return
	}

	h.emitter.Stop(customerID, webhookID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Event emitter stopped")
}

package main

import (
	"log"
	"net/http"
)

func main() {
	db := NewWebhookDB()
	client := NewRestClient(&http.Client{}, 3, 1)
	emitter := NewEmitter(db, client, 5)
	service := NewService(db)
	handler := NewHandler(service, emitter)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	http.HandleFunc("/register", handler.RegisterWebhook)
	http.HandleFunc("/", handler.GetWebhooks)
	http.HandleFunc("/emitter/start", handler.StartEmitter)
	http.HandleFunc("/emitter/stop", handler.StopEmitter)

	log.Println("sender service is running on port:8080")
	http.ListenAndServe(":8080", nil)
}

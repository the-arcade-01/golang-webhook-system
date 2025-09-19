package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/time/rate"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found or couldn't load it: %v", err)
	}

	limiter := rate.NewLimiter(5, 10)
	client := NewKafkaClient(os.Getenv("KAFKA_ADDR"), os.Getenv("KAFKA_TOPIC"))
	handler := NewHandler(client)

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	http.Handle("/webhook", ChainMiddleware(
		http.HandlerFunc(handler.RequestHandler),
		authMiddleware,
		func(h http.Handler) http.Handler { return rateLimitMiddleware(h, limiter) },
	))

	log.Println("webhook handler running on server :8081")
	http.ListenAndServe(":8081", nil)
}

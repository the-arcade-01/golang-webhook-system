package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file, %s", err)
	}

	db := NewEventsDB()
	client := NewKafkaClient(os.Getenv("KAFKA_ADDR"), os.Getenv("KAFKA_TOPIC"), os.Getenv("KAFKA_GROUP_ID"), db)
	go client.ConsumeEvents(context.Background())

	log.Println("event consumer running on port:8082")
	http.ListenAndServe(":8082", nil)
}

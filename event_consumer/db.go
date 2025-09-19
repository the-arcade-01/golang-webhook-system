package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type EventsDB struct {
	pool *pgxpool.Pool
}

func NewEventsDB() *EventsDB {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found or couldn't load it: %v", err)
	}

	cfg, err := pgxpool.ParseConfig(fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB")))
	if err != nil {
		log.Fatalf("Unable to parse DB config: %v", err)
	}

	cfg.MaxConns = 10
	cfg.MinConns = 2
	cfg.MaxConnIdleTime = 5 * time.Minute
	cfg.MaxConnLifetime = 30 * time.Minute

	pool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Fatalf("Unable to create DB pool: %v", err)
	}

	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping DB: %v", err)
	}

	log.Println("webhook db connection established")

	return &EventsDB{
		pool: pool,
	}
}

func (db *EventsDB) Insert(ctx context.Context, event Event) error {
	query := `
	INSERT INTO events (
    invoice_id,
    source_event_id,
    source_timestamp,
    event_type
	)
	VALUES ($1, $2, $3, $4)
	ON CONFLICT (invoice_id) DO UPDATE
	SET 
		source_event_id   = EXCLUDED.source_event_id,
		source_timestamp  = EXCLUDED.source_timestamp,
		event_type        = EXCLUDED.event_type,
		created_at        = NOW()
	WHERE EXCLUDED.source_timestamp > events.source_timestamp;
	`
	_, err := db.pool.Exec(ctx, query, event.Data.InvoiceID, event.EventID, event.Timestamp, event.EventType.ToString())
	return err
}

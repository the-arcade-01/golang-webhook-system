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

type WebhookDB struct {
	pool *pgxpool.Pool
}

func NewWebhookDB() *WebhookDB {
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

	return &WebhookDB{
		pool: pool,
	}
}

func (db *WebhookDB) Insert(ctx context.Context, customerID, webhookURL, secretToken string, webhookStatus WebhookStatus) (string, error) {
	query := `
		INSERT INTO webhooks (customer_id, webhook_url, secret_token, webhook_status)
		VALUES ($1, $2, $3, $4)
		RETURNING webhook_id;
	`
	var webhookID string
	err := db.pool.QueryRow(ctx, query, customerID, webhookURL, secretToken, webhookStatus.ToString()).Scan(&webhookID)
	return webhookID, err
}

func (db *WebhookDB) Get(ctx context.Context, customerID, webhookID string) (WebhookDetails, error) {
	query := `
		SELECT webhook_id, webhook_url, secret_token, webhook_status
		FROM webhooks
		WHERE customer_id = $1 AND webhook_id = $2;
	`
	var webhook WebhookDetails

	err := db.pool.QueryRow(ctx, query, customerID, webhookID).Scan(&webhook.WebhookID, &webhook.WebhookURL, &webhook.SecretToken, &webhook.WebhookStatus)
	return webhook, err
}

func (db *WebhookDB) GetAll(ctx context.Context) ([]WebhookDetails, error) {
	query := `
		SELECT webhook_id, webhook_url, webhook_status, secret_token
		FROM webhooks
	`
	rows, err := db.pool.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var webhooks []WebhookDetails
	for rows.Next() {
		var webhook WebhookDetails
		if err := rows.Scan(&webhook.WebhookID,
			&webhook.WebhookURL,
			&webhook.WebhookStatus,
			&webhook.SecretToken,
		); err != nil {
			return nil, err
		}
		webhooks = append(webhooks, webhook)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return webhooks, nil
}

func (db *WebhookDB) UpdateWebhookStatus(ctx context.Context, webhookID string, status WebhookStatus) error {
	query := `
		UPDATE webhooks
		SET webhook_status = $1
		WHERE webhook_id = $2;
	`
	_, err := db.pool.Exec(ctx, query, status.ToString(), webhookID)
	return err
}

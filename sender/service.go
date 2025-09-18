package main

import (
	"context"
	"errors"
	"log"
)

type Service struct {
	db *WebhookDB
}

func NewService(db *WebhookDB) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) RegisterWebhook(ctx context.Context, body WebhookRegisterBody) (*WebhookDetails, error) {
	secret := GenerateSecretKey(body.CustomerID, body.WebhookURL)
	webhookID, err := s.db.Insert(ctx, body.CustomerID, body.WebhookURL, secret, ACTIVE)

	if err != nil {
		log.Printf("error occurred while registering webhook, %s", err)
		return nil, errors.New("please try again later")
	}
	return &WebhookDetails{
		WebhookID:     webhookID,
		WebhookURL:    body.WebhookURL,
		SecretToken:   secret,
		WebhookStatus: ACTIVE,
	}, nil
}

func (s *Service) GetAll(ctx context.Context) ([]WebhookDetails, error) {
	webhooks, err := s.db.GetAll(ctx)
	if err != nil {
		log.Printf("error occurred while fetching all webhooks, %s", err)
		return nil, errors.New("please try again later")
	}
	return webhooks, nil
}

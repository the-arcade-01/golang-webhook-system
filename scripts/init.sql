\c webhooksdb;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE IF NOT EXISTS webhooks (
    webhook_id UUID PRIMARY KEY DEFAULT gen_random_uuid(), 
    customer_id TEXT NOT NULL,
    webhook_url TEXT NOT NULL,
    secret_token TEXT NOT NULL,
    webhook_status VARCHAR(20) NOT NULL CHECK (webhook_status IN ('ACTIVE', 'DISABLED')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS events (
    invoice_id VARCHAR(255) PRIMARY KEY,
    source_event_id VARCHAR(255) NOT NULL,
    source_timestamp TIMESTAMP NOT NULL,
    event_type VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
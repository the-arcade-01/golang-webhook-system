## Webhook System Prototype in Golang, Kafka, PostgreSQL

Read the blog for more implementation details: [TBA]

### Demo

https://github.com/user-attachments/assets/fdb24fe0-6670-4321-8af2-c7d5fcd21b44

### Tech Stack

- **Golang**: For all services.
- **Kafka**: Message broker for real-time data streaming.
- **PostgreSQL**: Persistent storage for webhook registrations and events.
- **Docker & Docker Compose**: Containerization and orchestration.
- **Nginx**: Reverse proxy and load balancing for webhook handler service.
- **REST API**: Service communication.
- **HMAC**: Request signing and verification for security.

### High Level Design

![HLD](./docs/HLD.png)

### Low Level Design (LLD)

This section describes the main components and logic implemented in the webhook system:

- **Sender Service**

  - Exposes REST APIs to register webhooks, start/stop event emitters, and view registered webhooks.
  - Generates and signs events using HMAC for security.
  - Emits events to Kafka for downstream processing.
  - Handles event simulation and manages emitter lifecycle per webhook.

- **Webhook Handler Service**

  - Receives webhook events via HTTP POST requests.
  - Verifies request authenticity using HMAC signature middleware.
  - Applies rate limiting and authentication middleware to protect endpoints.
  - Publishes received events to Kafka for further processing.

- **Event Consumer Service**

  - Consumes events from Kafka topic.
  - Processes and stores events in PostgreSQL database.
  - Implements retry logic for failed events; unprocessed events are retried and, after exceeding max retries, moved to a Dead Letter Queue (DLQ).

- **Database Layer**

  - PostgreSQL stores webhook registrations and event data.
  - Provides queries for registering, updating, and retrieving webhooks and events.

- **Kafka Integration**

  - Used for reliable, scalable event delivery between services.
  - Supports event retries and DLQ for failed events.

- **Security & Middleware**

  - HMAC-based request signing and verification for webhook authenticity.
  - Rate limiting and authentication middleware for API protection.

- **Docker & Orchestration**
  - All services are containerized using Docker.
  - Docker Compose orchestrates multi-service setup, scaling, and networking.

This design ensures secure, reliable, and scalable webhook delivery and processing with robust error handling and retry mechanisms.

### Run this project

#### 1. Create `.env` file in the root folder

```sh
cd golang-webhook-system
touch .env
```

#### 2. Update the `.env` variables

```sh
# Kafka
KAFKA_BROKER_ID=1
KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181
KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://kafka:29092,PLAINTEXT_HOST://localhost:9092
KAFKA_LISTENER_SECURITY_PROTOCOL_MAP=PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
KAFKA_INTER_BROKER_LISTENER_NAME=PLAINTEXT
KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1

# App configs
KAFKA_ADDR=kafka:29092
KAFKA_TOPIC=golang.webhook.system
KAFKA_GROUP_ID=webhook.consumer.group

# Postgres
POSTGRES_USER=<user>
POSTGRES_PASSWORD=<password>
POSTGRES_DB=webhooksdb
POSTGRES_PORT=5432
POSTGRES_HOST=postgres

# Webhook Secret Token
WEBHOOK_SECRET_TOKEN=<secret_token>
```

#### 3. Start core infrastructure services

```sh
docker compose up --build -d zookeeper kafka postgres
```

#### 4. Start sender service

```sh
docker compose up --build -d sender
```

#### 5. Register a webhook

```sh
curl -X POST http://localhost:8080/register -d '{ "customer_id": "ben@webhooksystem", "webhook_url": "http://webhook_handler:8081/webhook"}'
```

#### 6. Start remaining services

```sh
docker compose up --build -d webhook_handler nginx event_consumer
```

#### 7. Start event simulation

```sh
curl -X POST "http://localhost:8080/emitter/start?customer_id=ben@webhooksystem&webhook_id=d202ee17-9971-43a8-91f6-230cf5afd7e8"
```

#### 8. Stop event simulation

```sh
curl -X POST "http://localhost:8080/emitter/stop?customer_id=ben@webhooksystem&webhook_id=d202ee17-9971-43a8-91f6-230cf5afd7e8"
```

#### 9. Event retry & DLQ

Failed events are retried and pushed into Kafka again with a `retry_count` header. When `retry_count` exceeds the max, the event is moved to the DLQ (dead letter queue).

### API Reference

---

## Sender APIs

```sh
curl -X GET http://localhost:8080/ping
curl -X GET http://localhost:8080/
curl -X POST http://localhost:8080/register -d '{ "customer_id": "ben@webhooksystem", "webhook_url": "http://webhook_handler:8081/webhook"}'
curl -X POST "http://localhost:8080/emitter/start?customer_id=ben@webhooksystem&webhook_id=d202ee17-9971-43a8-91f6-230cf5afd7e8"
curl -X POST "http://localhost:8080/emitter/stop?customer_id=ben@webhooksystem&webhook_id=d202ee17-9971-43a8-91f6-230cf5afd7e8"
```

## Webhook Handler APIs

```sh
curl -X GET http://localhost:8081/ping
docker compose up --scale webhook_handler=3
docker exec -it <container_id> kafka-console-consumer --bootstrap-server localhost:9092 --topic golang.webhook.system --from-beginning
```

## Kafka CMDs

```sh
docker exec -it <container_id> kafka-console-consumer --bootstrap-server localhost:9092 --topic golang.webhook.system --from-beginning
```

## Docker CMDs

```sh
docker compose up --build -d zookeeper kafka postgres
docker compose up --build -d sender
docker compose up --build -d webhook_handler nginx event_consumer
```

## DB CMDs

```sh
psql -U root -d webhooksdb
\dt
select * from webhooks;
select * from events;
```

### References

1. https://pyemma.github.io/How-to-Design-Webhook/#security
2. https://beeceptor.com/docs/webhook-feature-design/#delivering-webhooks-at-scale-key-tech-design-considerations
3. https://systemdesignschool.io/problems/webhook/solution
4. https://tianpan.co/notes/166-designing-payment-webhook

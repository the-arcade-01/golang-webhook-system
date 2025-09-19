FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY event_consumer/go.mod event_consumer/go.sum ./

RUN go mod download
COPY event_consumer/ .
RUN go build -o event_consumer 

# STAGE - 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/event_consumer .
COPY .env .

EXPOSE 8082

CMD ["./event_consumer"]
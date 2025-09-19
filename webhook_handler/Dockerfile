FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY webhook_handler/go.mod webhook_handler/go.sum ./

RUN go mod download
COPY webhook_handler/ .
RUN go build -o webhook_handler 

# STAGE - 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/webhook_handler .
COPY .env .

EXPOSE 8081

CMD ["./webhook_handler"]
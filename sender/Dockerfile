FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY sender/go.mod sender/go.sum ./

RUN go mod download
COPY sender/ .
RUN go build -o sender 

# STAGE - 2
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/sender .
COPY .env .

EXPOSE 8080

CMD ["./sender"]
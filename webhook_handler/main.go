package main

import "fmt"

func main() {
	fmt.Println("Hello, Webhook handler")
}

/* TODO
1. API to receive the events
POST /webhook
Header: x-webhook-signature (hmac check middleware verification)
Body:
{
	event_id, event_type, timestamp, data: {invoice_id ...}
}
Response: after sending the event to kafka, send ACK back 200 status
if payload invalid send 400 back,
Also have Rate limit set to the API use proxy pattern before the service layer and after handler.

*/

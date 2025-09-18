package main

import "fmt"

func main() {
	fmt.Println("hello, event consumer")
}

/*
Get events from kafka, process them and puts them into Events DB,
also have logic for already processed events using event_id & timestamp check
also have out-of-orders events problem handling.

Intentionally fails some events so that they are moved back to kafka and then to DLQ.
*/

package main

import (
	"github.com/nats-io/go-nats"
	"log"
)

func main() {
	con, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Println("connect server failed")
		return
	}
	defer con.Close()
	log.Println("Connected to " + nats.DefaultURL)
	// Publish message on subject
	subject := "foo"
	con.Publish(subject, []byte("Hello NATS"))
	log.Println("Published message on subject " + subject)
	con.Close()
}

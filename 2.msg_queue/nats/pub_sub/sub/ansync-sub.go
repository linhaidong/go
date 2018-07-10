package main

import (
	"github.com/nats-io/go-nats"
	"log"
	"runtime"
)

func main() {
	con, _ := nats.Connect(nats.DefaultURL)
	log.Println("Connect to" + nats.DefaultURL)

	log.Printf("subscribing to subject 'foo'\n")
	con.Subscribe("foo", func(msg *nats.Msg) {
		log.Printf("recvive message:%s\n", string(msg.Data))
	})
	runtime.Goexit()
}

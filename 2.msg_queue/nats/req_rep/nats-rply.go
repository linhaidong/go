package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/go-nats"
)

// NOTE: Use tls scheme for TLS, e.g. nats-rply -s tls://demo.nats.io:4443 foo hello
func usage() {
	log.Fatalf("Usage: nats-rply [-s server][-t] <subject> <response>\n")
}

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var showTime = flag.Bool("t", false, "Display timestamps")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 2 {
		usage()
	}

	nc, err := nats.Connect(*urls)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}

	//sub 订阅的消息, reply 返回的消息
	subj, reply, i := args[0], args[1], 0

	//订阅消息
	nc.Subscribe(subj, func(msg *nats.Msg) {
		i++
		printMsg(msg, i)
		//返回消息
		fmt.Printf("call back fun\r\n")
		nc.Publish(msg.Reply, []byte(reply))
	})
	nc.Flush()

	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on [%s]\n", subj)
	if *showTime {
		log.SetFlags(log.LstdFlags)
	}

	runtime.Goexit()
}

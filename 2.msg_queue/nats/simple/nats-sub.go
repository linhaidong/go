package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/go-nats"
	"log"
	"runtime"
	"time"
)

// NOTE: Use tls scheme for TLS, e.g. nats-sub -s tls://demo.nats.io:4443 foo
func usage() {
	log.Fatalf("Usage: nats-sub [-s server] [-t] <subject> \n")
}

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var showTime = flag.Bool("t", false, "Display timestamps")
	var block = flag.Bool("b", true, "sync subscribe")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		usage()
	}
	if *showTime {
		log.SetFlags(log.LstdFlags)
	}

	//链接到server
	nc, err := nats.Connect(*urls)
	if err != nil {
		log.Fatalf("Can't connect: %v\n", err)
	}

	subj, i := args[0], 0
	//同步消息
	if *block {
		fmt.Println(subj)
		sub, err := nc.SubscribeSync(subj)
		if err != nil {
			fmt.Println("sub message error", err.Error())
			return
		}
		msg, err := sub.NextMsg(5 * time.Second)
		if err != nil {
			fmt.Println("recv msg", err.Error())
			return
		}
		fmt.Println("get msg is ", string(msg.Data))
		sub.Unsubscribe()
		nc.Close()
	} else {
		//异步消息
		//订阅主题，所有的订阅都可以收到消息
		nc.Subscribe(subj, func(msg *nats.Msg) {
			i += 1
			printMsg(msg, i)
		})
		//Flush will perform a round trip to the server and return when it receives the internal reply.
		nc.Flush()
	}

	log.Printf("Listening on [%s]\n", subj)
	/*
		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		}

			go func() {
				fmt.Println("go route exit 1111!!!")
				runtime.Goexit()
				fmt.Println("go route exit 222!!!")
			}()
						Calling Goexit from the main goroutine terminates that goroutine without func main returning.
				        Since func main has not returned, the program continues execution of other goroutines. If all other goroutines exit, the program crashes.
	*/
	runtime.Goexit()
}

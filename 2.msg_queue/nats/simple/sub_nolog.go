package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/go-nats"
	"runtime"
	"time"
)

func printMsg(m *nats.Msg, i int) {
	fmt.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
}

func main() {
	var urls = flag.String("s", nats.DefaultURL, "The nats server URLs (separated by comma)")
	var block = flag.Bool("b", true, "sync subscribe")

	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("param small")
	}

	//链接到server
	nc, err := nats.Connect(*urls)
	if err != nil {
		fmt.Printf("Can't connect: %v\n", err)
		return
	}

	subj, i := args[0], 0
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
		//订阅主题，所有的订阅都可以收到消息
		fmt.Println("use sub by func")
		nc.Subscribe(subj, func(msg *nats.Msg) {
			i += 1
			printMsg(msg, i)
		})
		//Flush will perform a round trip to the server and return when it receives the internal reply.
		nc.Flush()
	}

	fmt.Printf("Listening on [%s]\n", subj)
	/*
		if err := nc.LastError(); err != nil {
			fmt.Println(err.Error())
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

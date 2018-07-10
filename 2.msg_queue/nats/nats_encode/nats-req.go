package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	//NewEncodedConn will wrap an existing Connection and utilize the appropriate registered encoder.
	c, _ := nats.NewEncodedConn(nc, "json")
	defer c.Close()
	type person struct {
		Name    string
		Address string
		Age     int
	}

	ch := make(chan *person)
	//订阅主题和chan绑定
	c.BindRecvChan("hello", ch)

	me := &person{Name: "derek", Age: 22, Address: "85 Second St"}
	c.Publish("hello", me)

	// Receive the publish directly on a channel
	who := <-ch

	fmt.Printf("%v , age %d, says %s!\n", who.Name, who.Age, who.Address)
}

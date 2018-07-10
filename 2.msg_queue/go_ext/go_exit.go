package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"
)

type message struct {
	normal bool                   //true means exit normal, otherwise
	state  map[string]interface{} //goroutine state
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	mess := make(chan message, 10)
	for i := 0; i < 100; i++ {
		go worker(mess)
	}
	supervisor(mess)
}

func worker(mess chan message) {
	defer func() {
		exit_message := message{state: make(map[string]interface{})}
		i := recover()
		if i != nil {
			exit_message.normal = false
		} else {
			exit_message.normal = true
		}
		mess <- exit_message
	}()
	now := time.Now()
	seed := now.UnixNano()
	rand.Seed(seed)
	num := rand.Int63()
	if num%2 != 0 {
		panic("not evening")
	} else {
		runtime.Goexit()
	}
}

func supervisor(mess chan message) {
	for i := 0; i < 100; i++ {
		m := <-mess
		switch m.normal {
		case true:
			log.Println(i, "exit normal, nothing serious!")
		case false:
			log.Println(i, "exit abnormal, something went wrong")
		}

	}
}

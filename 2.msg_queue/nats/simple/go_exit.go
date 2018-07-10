package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("begin")
	go func() {
		time.Sleep(1 * time.Second)
	}()
	runtime.Goexit()
}

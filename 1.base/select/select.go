package main

import (
	"fmt"
	"time"
)

/* --------------------------------------------------------------------------*/
/**
* @Synopsis  判断缓冲区是否满了
*
* @Returns
 */
/* ----------------------------------------------------------------------------*/
func test1() {
	ch := make(chan int, 1)
	ch <- 1
	select {
	case ch <- 2:
		fmt.Println("input messge success")
	default:
		fmt.Println("channedl is full")
	}
	//i := <-ch
	//fmt.Printf("get message %d\r\n", i)
}
func test2() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	//select 操作具有时效性,当时判断不成功便返回
	select {
	case <-ch1:
		fmt.Println("ch1 pop one element")
	case <-ch2:
		fmt.Println("ch2 pop one element")
	default:
		fmt.Println("default")
	}
}

func test3() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()
	//select 操作具有时效性,加上default, 判断不成功便返回
	//去掉default, 会一直阻塞
	select {
	case <-ch1:
		fmt.Println("ch1 pop one element")
	case <-ch2:
		fmt.Println("ch2 pop one element")
		//default:
		//	fmt.Println("default")
	}
}
func main() {
	test1()
}

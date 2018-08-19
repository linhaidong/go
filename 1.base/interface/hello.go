package main

import (
	"fmt"
)

//定义接口类型
type ISayHello interface {
	SayHello()
}

type Person struct{}

//实现接口
func (person Person) SayHello() {
	fmt.Printf("Hello!\n")
}

type Duck struct{}

func (duck Duck) SayHello() {
	fmt.Printf("ga ga ga!\n")
}

func greeting(i ISayHello) {
	i.SayHello()
}

func test1() {
	person := Person{}
	duck := Duck{}
	var i ISayHello
	i = person
	//逻辑执行
	greeting(i)
	//
	person.SayHello()

	i = duck
	greeting(i)
}

/*interface 是一个类型*/
func test2() {
	//meth为interface类型的变量
	var meth interface {
		SayHello()
	}
	var d1 Duck
	meth = d1
	meth.SayHello()
}

/*断言, inteface{} 转换为自有类型*/
func test3() {
	var i interface{}
	i = 3
	a := i.(int)
	fmt.Printf("a = %d\n", a)
}

package main

import "fmt"

/*define interface*/
type phone interface {
	call()
}

type interface1 interface {
	iface_fun1()
}

type interface2 interface {
	call()
}

type iphone struct {
}
type sphone struct {
}

/*struct define interface*/
func (name iphone) call() {
	fmt.Println("i am iphone")
}

func (name iphone) iface_fun1() {
	fmt.Println("iphone interface 1 function")
}

func (name sphone) call() {
	fmt.Println("i am sphone")
}

func call() {
	fmt.Println("i'm define call func'")
}

/******************************************************************/
/***********              函数接口使用     ************************/
/******************************************************************/
type adder interface {
	add(a int)
}
type Handerfunc func(a int)

func (f Handerfunc) add(a int) {
	fmt.Println("call func")
	a += 100
	f(a)
}

func add_fun(a int) {
	fmt.Printf("a = %d\n", a)
}

func func_interface() {
	a := 100
	var b, c int = 100, 200

	var handler adder
	//函数强制转换，为接口类型
	handler = Handerfunc(add_fun)
	//调用接口的实现函数
	handler.add(a)
	handler.add(b)
	handler.add(c)
}

func struct_interface() {
	call()
	var ips iphone
	var sps sphone

	var pi phone
	var i1 interface1
	var i2 interface2

	/*interface val to struct*/
	//同样的接口，指向不同的实现
	pi = ips
	pi.call()

	pi = sps
	pi.call()

	//define interface val
	var i_var phone
	//interface val eq struct
	i_var = new(iphone)
	//use interface val call struct function
	i_var.call()

	i1 = ips
	i1.iface_fun1()
	//接口之间赋值
	fmt.Println("assign interface to another")
	i2 = pi
	i2.call()
}

/******************************************************************/
/***********            指针作为接收器     ************************/
/******************************************************************/
type caller interface {
	caller()
}
type another struct {
}

func (a another) caller() {
	fmt.Println("another normal caller")
}

func (p *iphone) caller() {
	fmt.Println("iphone point caller")
}

func point_recver() {
	var c caller
	var j iphone
	//invalid indirect of j (type iphone)
	//(*j).caller()
	//编译器会隐式地帮我们用&p去调用caller这个方法。这种简写方法只适用于“变量”
	j.caller()
	(&j).caller()
	/*iphone does not implement caller (caller method has pointer receiver)*/
	c = &j
	c.caller()

	i := &iphone{}
	i.caller()
	c = i
	c.caller()

	a := another{}
	c = a
	a.caller()
}

package main

import (
	"fmt"
	"math"
)

/*GO  语言函数的声明形式：
func fun_name(arg1 type, arge type) 返回值列表（可为多个返回值）
*/
func max(num1 int, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
func swap(x string, y string) (string, string) {
	return y, x
}

/*引用传递，形参声明形参后面加*， 调用时变量前加&*/
func int_swap(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}

/*闭包函数：函数作为返回值，
  声明时，函数返回值前添加 func()关键字
  使用时直接使用函数内的变量
*/
func get_seq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/*为结构体声明方法，相当于类的方法
声明方法：函数名称见面添加（类名称  类类型）*/
/*声明结构体*/
type Circle struct {
	radius float64
}

/*声明结构体方法*/
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
func main() {
	var a int = 100
	var b int = 200
	var x string = "hello"
	var y string = "world"
	var ret int
	ret = max(a, b)
	fmt.Printf("max is %d\r\n", ret)
	m, n := swap(x, y)
	println(x, y, m, n)
	int_swap(&a, &b)
	fmt.Printf("a = %d, b = %d\r\n", a, b)
	/*函数变量，相当于函数指针,不用指定函数名称，只需要声明func关键字
	  函数的名称用func关键字代替*/
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquareRoot(9))

	nextNumer := get_seq()
	fmt.Println(nextNumer())
	fmt.Println(nextNumer())
	fmt.Println(nextNumer())

	nextNumer1 := get_seq()
	fmt.Println(nextNumer1())
	fmt.Println(nextNumer1())

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of Circle = ", c1.getArea())

	var point *int
	point = &b
	fmt.Println(*point, b)
}

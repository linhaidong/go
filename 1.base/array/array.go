package main

import "fmt"

/*向函数传递数组*/
/*函数声明的数组[]中可带数字也可不带数字
在函数使用的时候，要传递同样定义的数组
如果函数声明带数字，函数使用的数组变量，定义时必须带相同的数组
不带数字的函数声明的用法相同*/
func get_average(array [5]int, num int) float32 {
	var i, sum int
	var avg float32

	for i = 0; i < num; i++ {
		sum += array[i]
	}

	avg = float32(sum / num)
	return avg
}

func base_test() {
	/*定义数组, 类型前面添加[]即为定义数组*/
	var array1 [5]int
	//定义长度为100的数组，前99为0，最后一位为-1
	r := [...]int{99: -1}
	/*直接定义数组变量, 必须定义与数组形参函数相同的数组类型*/
	var array2 = [5]int{1, 2, 3, 4, 5}
	/*定义多维数组*/
	var num_array [2][2]int
	var num_array2 = [2][2]int{{2, 2}, {3, 3}}

	/*数组int类型默认初始化为0*/
	var i, j int
	for i = 0; i < 5; i++ {
		fmt.Println(array1[i])
	}
	for i = 0; i < 5; i++ {
		fmt.Println(array2[i])
	}
	for i = 0; i < 2; i++ {
		for j = 0; j < 2; j++ {
			/*输出单列*/
			fmt.Println(num_array[i])
			fmt.Println(num_array[i][j])
		}
	}
	for i = 0; i < 2; i++ {
		for j = 0; j < 2; j++ {
			fmt.Println(num_array2[i])
			fmt.Println(num_array2[i][j])
		}
	}

	fmt.Println("array average is ", get_average(array2, 5))
}

/**********************************************************/
type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func string_array() {

	symbol := [...]string{USD: "$", EUR: "A", GBP: "3", RMB: "Y"}
	fmt.Println(RMB, symbol[RMB])
}

func array_compare() {
	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	//d := [3]int{1, 2}
	//fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int
}

func main() {
	base_test()
	string_array()
}

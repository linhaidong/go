package main

import "fmt"

/* --------------------------------------------------------------------------*/
/**
* @Synopsis  空切片
*
* 一个零值的slice等于nil。一个nil值的slice并没有底层数组。
* 一个nil值的slice的长度和容量都是0，但是也有非nil值的slice的长度和容量也是0的，
* 例如[]int{}或make([]int, 3)[3:]。
*如果你需要测试一个slice是否是空的，使用len(s) == 0来判断，而不应该用s == nil来判断
* @Returns
 */
/* ----------------------------------------------------------------------------*/
func slice_empty() {
	fmt.Println("**************************************")
	/*定义空的切片，等于nil*/
	var slice []int
	/*切片在未初始化之前默认为nil，长度为0*/
	printslice(slice)
	/*if 后面必须要加{}否则会报错*/
	if slice == nil {
		fmt.Println("slice is empty")
	}
	if len(slice) == 0 {
		fmt.Println("slice is empty")
	}
}

func slice_init() {
	fmt.Println("**************************************")
	/*make 创造新的切片, 可以制定长度和容量*/
	var slice1 []int = make([]int, 3, 5)
	/*GO语言可以自动识别变量类型，切片初始化，并赋值*/
	slice2 := []int{1, 2, 3}
	var slice10 []int = []int{1, 2, 3}
	var array = [5]int{1, 2, 3, 4, 5}
	slice3 := array[0:4]
	slice4 := slice1[:]
	slice5 := make([]int, 3, 5)
	printslice(slice1)
	printslice(slice2)
	printslice(slice3)
	printslice(slice4)
	printslice(slice5)
	fmt.Printf("slice10:")
	printslice(slice10)

	/*slice 的截取*/
	num := slice3[1:3]
	printslice(num)
}

func slice_append() int {
	fmt.Println("**************************************")
	var num []int
	printslice(num)
	/*append 函数向数组后面追加元素*/
	num = append(num, 0)
	printslice(num)

	num = append(num, 1)
	printslice(num)

	num = append(num, 2, 3, 4)
	printslice(num)

	num1 := make([]int, len(num), cap(num)*2)
	/*copy函数用于复制切片*/
	copy(num1, num)
	printslice(num1)
	return 0
}

func slice_operate() {
	fmt.Println("**************************************")
	var b = []int{1, 2, 3, 4}
	fmt.Printf("b  = %v\r\n", b)

	var slice []int
	//append 为内置的函数,专用于slice的添加
	//append函数返回值必须有变量接收，不然编译器会报错
	slice = append(slice, 0, 1, 2, 3, 4, 5)
	slice = append(slice, []int{6, 7, 8}...)
	fmt.Printf("slice lenth  %d\r\n", len(slice))
	fmt.Printf("slice is %v\r\n", slice)

	//del elemt
	index := 5
	fmt.Printf("slice addr 0x%p\r\n", &slice)
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Printf("slice addr 0x%p\r\n", &slice)
	fmt.Printf("after delete slice: %v\r\n", slice)

	//copy
	//复制只复制该slic的大小
	//make([]int, len, capalibty)
	csli := make([]int, 0, 10)
	copy(csli, slice)
	fmt.Printf("copy slice:%v\r\n", csli)

	csli2 := make([]int, 5, 10)
	copy(csli2, slice)
	fmt.Printf("copy slice:%v\r\n", csli2)
}

func printslice(x []int) {
	/*cap函数获取切片的容量，len函数获取切片的长度*/
	fmt.Printf("len = %d, cap =%d slice =%d\r\n", len(x), cap(x), x)
}

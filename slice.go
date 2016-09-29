package main

import "fmt"

func main(){
    /*定义空的切片，等于nil*/
    var slice [] int
    /*make 创造新的切片, 可以制定长度和容量*/
    var slice1 [] int = make([]int, 3, 5)
    /*GO语言可以自动识别变量类型，切片初始化，并赋值*/
    slice2 := [] int {1,2,3}
    var slice10 [] int = [] int {1,2,3} 
    var array  = [5] int {1,2,3,4,5}
    slice3 := array[0:4]
    slice4 := slice1[:]
    slice5 := make([] int, 3, 5)
    /*if 后面必须要加{}否则会报错*/
    if(slice == nil){
        fmt.Printf("slice is empty")
    }
    /*切片在未初始化之前默认为nil，长度为0*/
    printslice(slice)
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
    fmt.Printf("/***  test   ***/\r\n")
    test()
}
func test() int {
    var num [] int
    printslice(num)
    /*append 函数向数组后面追加元素*/ 
    num = append(num, 0)
    printslice(num)
    
   
    num = append(num, 1)
    printslice(num)

    num = append(num, 2,3,4)
    printslice(num)

    num1 := make([] int, len(num), cap(num)*2)
    /*copy函数用于复制切片*/
    copy(num1, num)
    printslice(num1)
    return 0
}
func printslice(x []int){
    /*cap函数获取切片的容量，len函数获取切片的长度*/
    fmt.Printf("len = %d, cap =%d slice =%d\r\n", len(x), cap(x), x)
}

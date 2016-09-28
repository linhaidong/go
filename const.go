package main

import "fmt"
import "unsafe"
const (
    aa = "abc"
    bb = len(aa)
    cc = unsafe.Sizeof(aa)
)
const (
    a1 = iota
    a2
    a3
)
const (
    b1 = iota
    b2
    b3
)
const (
    i = 1<<iota
    j = 3<<iota
   // i = 1
   // j = 6
    k
    l
)
func main() {
    const LEGTH int = 10
    const WIDTH int = 5
    const a , b ,c = 1, false , "str"
    //fmt.Printf("iota is %d", iota)
    area := LEGTH*WIDTH
    fmt.Printf("area:%d", area)
    println()
    println(a, b, c)
    println(aa, bb, cc)
    println(a1, a2, a3)
    println(b1, b2, b3)
    println(i,j,k,l)
}

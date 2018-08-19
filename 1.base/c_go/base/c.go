package main

// #include <stdio.h>
// #include <stdlib.h>
/*
void print(char *str) {
    printf("%s\n", str);
}
*/
import "C"

//import "C" 和上面的注释不能有空格,否则会出现编译错误
//could not determine kind of name for C.test
import "unsafe"

func main() {
	s := "Hello Cgo"
	cs := C.CString(s)
	C.print(cs)
	C.free(unsafe.Pointer(cs))
}

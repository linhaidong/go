package main

import "fmt"

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

/* --------------------------------------------------------------------------*/
/**
* @Synopsis  数组翻转
*
* @Returns
 */
/* ----------------------------------------------------------------------------*/
func Reverse_array() {
	a := []int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)
}

/* --------------------------------------------------------------------------*/
/**
* @Synopsis  数组左移两位
*
* @Returns
 */
/* ----------------------------------------------------------------------------*/
func Rotate_array() {
	s := []int{0, 1, 2, 3, 4, 5}
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s)
}

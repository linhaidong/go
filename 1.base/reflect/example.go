// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

/* --------------------------------------------------------------------------*/
/**
* @Synopsis  利用value实现函数模板, 为函数变量赋值, 进行不同的操作
*
* @Returns
 */
/* ----------------------------------------------------------------------------*/
func create_function() {
	// swap is the implementation passed to MakeFunc.
	// It must work in terms of reflect.Values so that it is possible
	// to write code without knowing beforehand what the types
	// will be.
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	// makeSwap expects fptr to be a pointer to a nil function.
	// It sets that pointer to a new function created with MakeFunc.
	// When the function is invoked, reflect turns the arguments
	// into Values, calls swap, and then turns swap's result slice
	// into the values returned by the new function.
	//函数变量赋值
	makeSwap := func(fptr interface{}) {
		// fptr is a pointer to a function.
		// Obtain the function value itself (likely nil) as a reflect.Value
		// so that we can query its type and then set the value.
		//ValueOf returns a new Value initialized to the concrete value stored in the interface i.
		//获取真正的类型
		fn := reflect.ValueOf(fptr).Elem()

		// Make a function of the right type.
		//创建真正的函数, 相当与函数模板实例化
		v := reflect.MakeFunc(fn.Type(), swap)

		// Assign it to the value fn represents.
		fn.Set(v)
	}

	// Make and call a swap function for ints.
	//定义函数变量
	var intSwap func(int, int) (int, int)
	if intSwap == nil {
		fmt.Println("init swap is null")
	}
	//根据函数类型为,函数变量赋值
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	// Make and call a swap function for float64s.
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))

	// Output:
	// 1 0
	// 3.14 2.72
}

func ExampleStructTag_Lookup() {
	type S struct {
		F0 string `alias:"field_0"`
		F1 string `alias:""`
		F2 string
	}

	s := S{}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if alias, ok := field.Tag.Lookup("alias"); ok {
			if alias == "" {
				fmt.Println("(blank)")
			} else {
				fmt.Println(alias)
			}
		} else {
			fmt.Println("(not specified)")
		}
	}

	// Output:
	// field_0
	// (blank)
	// (not specified)
}

func ExampleTypeOf() {
	// As interface types are only used for static typing, a
	// common idiom to find the reflection Type for an interface
	// type Foo is to use a *Foo value.
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))

	// Output:
	// true
}

/* --------------------------------------------------------------------------*/
/**
* @Synopsis  动态的创建对象
             首先StructOf创建类型,
*
* @Returns
*/
/* ----------------------------------------------------------------------------*/
func create_struct() {
	//StructOf函数,创建struct type, 返回Type
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})

	//reflect.New()根据类型,分配相应的空间,返回相应的指针
	//返回的value为指针类型的,如果要
	//New returns a Value representing a pointer to a new zero
	//value for the specified type. That is, the returned Value's Type is PtrTo(typ).
	//New函数,返回相应类型的指针
	//Elem returns the value that the interface v contains or that the pointer v points to
	//Elem 相当与对"*",相当与对指针取值,
	//Elem 返回value类型
	v := reflect.New(typ).Elem()

	//Field returns the i'th field of the struct v.
	v.Field(0).SetFloat(0.4)
	v.Field(1).SetInt(2)
	//Addr returns a pointer value representing the address of v.
	//Addr相当于&
	//Interface returns v's current value as an interface{}.
	//Interface 把value转换为interface{}
	s := v.Addr().Interface()

	fmt.Printf("value: %+v\n", s)

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}

	fmt.Printf("json:  %s", w.Bytes())

	//把josn串转换到生成的结构体中
	r := bytes.NewReader([]byte(`{"height":1.5,"age":10}`))
	if err := json.NewDecoder(r).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)

	// Output:
	// value: &{Height:0.4 Age:2}
	// json:  {"height":0.4,"age":2}
	// value: &{Height:1.5 Age:10}
}

/*
func main() {
	//ExampleStructOf()
	ExampleMakeFunc()
}*/

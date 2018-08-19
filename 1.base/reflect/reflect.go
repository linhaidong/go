package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func base_use() {
	var a int = 5
	t := reflect.TypeOf(a)
	fmt.Println("type:", t)

	//reflect.TypeOf 返回的是一个动态类型的接口值, 它总是返回具体的类型
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))
	//输出*os.File
	//fmt %T 参数, 内部使用 reflect.TypeOf 来输出
	fmt.Printf("%T\n", 3)
	fmt.Printf("%T\n", w)

	//返回类型的value类型
	v := reflect.ValueOf(3)
	fmt.Printf("%v\n", v)
	fmt.Println(v.String())
	t = v.Type()
	fmt.Println(t.String())

	//获取value类型的值，返回interface{}
	x := v.Interface()
	i := x.(int)
	fmt.Printf("%d\n", i)
}

/* --------------------------------------------------------------------------*/
/**
* @Synopsis  结构体的tag操作, 类型操作
            获得结构体内部的具体的类型的信息
*
* @Param )
*
* @Returns
*/
/* ----------------------------------------------------------------------------*/
func struct_tag_oper() {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	//TypeOf接受interface{}作为参数,表示可以接收任意类型
	//返回Type接口,接口中有elem, Field等方法可以调用
	st := reflect.TypeOf(s)
	//Field 为Type接口的方法,参数为tag的序号, 返回一个structField
	//表示一个具体的结构体的类型

	//structField结构,用标准的形式,来标识一个结构体的字段
	// Field returns a struct type's i'th field.
	field := st.Field(0)
	//A StructTag is the tag string in a struct field.
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

	// Output:
	// blue gopher
}

/* --------------------------------------------------------------------------*/
/**
* @Synopsis  通用的取值计算方式,
*
* @Param interface{
}
*
* @Returns
*/
/* ----------------------------------------------------------------------------*/
func get_name(val interface{}) {
	lo := reflect.ValueOf(val).Elem()
	//获得结构体的Field(结构体的一项)
	m2 := lo.FieldByName("Name")
	fmt.Println(m2.Addr())
	if m2.CanAddr() {
		fmt.Println("value can addr")
	}
	if m2.CanSet() {
		fmt.Println("value can set")
	}
	fmt.Println(m2.Kind())
	//m2.Set(reflect.ValueOf("haha"))
	new_name := "jaja"
	//设置结构体name field的值
	m2.SetString(new_name)
}

func set_age(person interface{}) {
	p := reflect.ValueOf(person).Elem()
	age := p.FieldByName("Age")
	age.SetInt(1000)
}
func test_name() {
	//结构体变量名称必须为大写
	type man struct {
		Name string
		Age  int
	}
	type woman struct {
		Name string
		Age  int
	}
	m := &man{Name: "jodan", Age: 10}
	w := &woman{Name: "mimi", Age: 100}
	get_name(m)
	get_name(w)
	//m.Name = "newname"
	fmt.Println("name:", m.Name)
	fmt.Println("name:", w.Name)
	set_age(m)
	set_age(w)
	fmt.Println("man age:", m.Age)
	fmt.Println("woman age:", w.Age)
}

func main() {
	base_use()
	struct_tag_oper()
	create_struct()
	create_function()
	test_name()
}

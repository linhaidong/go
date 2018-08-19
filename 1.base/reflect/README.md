反射:
反射机制的优点就是可以实现"动态创建对象和编译".
动态的创建对象
泛型的操作不同对象里

type:
主要以类型进行作为操作对象.
定义了structType, 里面包含结构体的相关的属性,只要按照结构体类型,定义相关的结构,
就可以动态的定义结构体对象.
    func ArrayOf(count int, elem Type) Type
    func ChanOf(dir ChanDir, t Type) Type
    func FuncOf(in, out []Type, variadic bool) Type
    func MapOf(key, elem Type) Type
    func PtrTo(t Type) Type
    func SliceOf(t Type) Type
    func StructOf(fields []StructField) Type
上面的函数均是创建相应的类型,比如StructOf创建结构体类型:
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





Value:
type Value struct {
	// typ holds the type of the value represented by a Value.
	typ *rtype

	// Pointer-valued data or, if flagIndir is set, pointer to data.
	// Valid when either flagIndir is set or typ.pointers() is true.
	ptr unsafe.Pointer

	// flag holds metadata about the value.
	flag
}


func ValueOf(i interface{}) Value
ValueOf returns a new Value initialized to the concrete value stored in the interface i.
ValueOf(nil) returns the zero Value.


elem函数,对指针value或interfce进行取值,相当与指针的*操作
func (v Value) Elem() Value
Elem returns the value that the interface v contains or that the pointer v points to.
It panics if v's Kind is not Interface or Ptr. 
It returns the zero Value if v is nil.turns the zero Value.


value 提供了类型值的通用操作, 创建,访问,取值, 长度等操作

type 表示类型, value用于对相应的基本类型进行操作
创建:
MakeChan, MakeFunc, MakeMap, MakeSlice用于创建相应类型的值,
New()函数根据type分配相应的空间并初始化, 创建相应的类型的实例 

reflect.Value 和 interface{} 都能装载任意的值. 所不同的是, 一个空的接口隐藏了值内部的表示方式和所有方法, 因此只有我们知道具体的动态类型才能使用类型断言来访问内部的值(就像上面那样), 内部值我们没法访问. 相比之下, 一个 Value 则有很多方法来检查其内容, 无论它的具体类型是什么. 

[参考网址]
https://golang.org/pkg/reflect/#Value.Set

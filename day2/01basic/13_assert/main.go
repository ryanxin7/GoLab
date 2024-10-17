package main

import (
	"fmt"
)


type Animal interface {
	Say() string
}

//Cat和Dog结构体都实现了这个接口方法，那么就实现了这个接口

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (c *Cat) Say() string {
	return c.Name + "喵喵喵"
}

func (d *Dog) Say() string {
	return d.Name + "汪汪汪"
}

func main() {
	c := &Cat{Name: "小白"}
	d := &Dog{Name: "大黑"}
	//把一个具体的结构体（比如 Cat 或 Dog）转换为接口类型（比如 Animal）
	var p1 Animal
	p1 = c // 这里 c 是 *Cat 类型，但被转换成了 Animal 接口, Go 自动把 *Cat 类型的 c 赋值给了接口类型的 p1。
	fmt.Println(p1.Say())
	p1 = d
	fmt.Println(p1.Say())
	//p1 作为 Animal 类型的变量，可以指向不同的结构体，并根据不同的结构体调用各自的 Say() 方法。
	fmt.Println(transData(c))
	//fmt.Println(transData(d))

	var x interface{}
	x = true
	v2, ok := x.(string)
	if ok {
		fmt.Println(v2)
		fmt.Printf("值:%v 类型:%T",v2, v2)
	} else {
		fmt.Println("非字符串类型")

	}

}

//类型转换举例
//结构体转换为interface，是从下往上转，是隐式转换
//interface转换为结构体，是从上往下转，是显示转换，专业名词：断言
//没断言之前，只能拿到接口的方法，断言后，能拿到对应类型的所有属性和方法

func transData(a Animal) string {
	v, ok := a.(*Cat)
	if ok {
		fmt.Println("断言成功")
	} else{
	fmt.Println("断言失败")
    }
	fmt.Println(v.Name)
	return fmt.Sprintf("%s%s", a.Say(), "处理后")

//Type assertion allows you to retrieve the original type from a value stored in an interface{}.
//类型断言允许你从存储在 interface{} 中的值中取回原始类型。

//In Go, when a struct or any other type is assigned to an interface{}, Go wraps the original value and its type inside the interface.
//在 Go 语言中，当一个结构体或其他类型赋值给 interface{} 时，Go 会将原始值和它的类型包装在接口中。

//The type assertion checks if the value stored in the interface matches a specific type.
//类型断言用于检查接口中存储的值是否与特定类型匹配。

//Without a type assertion, you can only access the methods defined by the interface, not the original value's fields or methods.
//没有类型断言时，你只能访问接口定义的方法，而不能访问原始值的字段或方法。


//value, ok := x.(OriginalType)

//value: The variable that stores the value of type OriginalType if the assertion is successful.
//value：如果断言成功，存储原始类型值的变量。
	//
	//ok: A boolean that indicates whether the assertion was successful (true) or failed (false).
	//ok：一个布尔值，指示断言是否成功（true）或失败（false）。


}


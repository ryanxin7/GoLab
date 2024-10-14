package main

//Animal 接口定义了行为
//Cat 和 Dog 实现了这个接口的行为
//通过多态，Animal 类型的变量可以表现为 Cat 或 Dog，并调用对应的 Say() 方法

import "fmt"

//结构体实现了接口中的所有方法，那么我们就说这个结构体实现了这个接口
//类型转换，结构体转接口，接口转结构体
//实现接口后，就能实现特定的功能，比如golang自定义排序
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
	fmt.Println(transData(d))
}

//类型转换举例
func transData(a Animal) string {
	return fmt.Sprintf("%s%s", a.Say(), "处理后")
}

//第一个 %s 被替换成了 "小白喵喵喵"（c.Say() 的返回值）
//第二个 %s 被替换成了 "处理后"
//所以最后的输出是 "小白喵喵喵处理后"。

//在编程中，“多态”就是指同一个接口，可以让不同的对象表现出不同的行为。

//就像你在代码里看到的，接口 Animal 有一个方法 Say()。不管是 Cat 还是 Dog，它们都实现了这个 Say() 方法，但是表现的行为不一样：

//Cat 的 Say() 输出“喵喵喵”
//Dog 的 Say() 输出“汪汪汪”
//当我们把 Cat 和 Dog 都当作 Animal 类型来处理时，虽然它们都属于同一个接口 Animal，但在调用 Say() 的时候，具体的行为由它们自己的实现决定。这就是“多态”！

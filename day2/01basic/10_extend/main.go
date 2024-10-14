package main

import "fmt"

//父结构
type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p *Person) SayHello() {
	fmt.Println(p)
	fmt.Println("this is form Person")

}

//子结构
type Student struct {
	School string
	//People *Person
	People Person
	//Person自定义类型
}

//一般情况下不使用golang继承，用嵌套结构体替代，因为嵌套结构体比较简单，易读

func main() {
	stu := &Student{
		School: "middle",
		People: Person{}, // 必须先为 People 分配内存
	}
	//模拟继承的关系，帮我们自动处理
	//stu.Name = stu.Person.Name
	//stu.SayHello() = stu.Person.SayHello()
	//属性和方法都能继承
	stu.People.Name = "Leo"
	stu.People.Age = 30
	stu.People.Sex = "Male"
	fmt.Println(stu)
}

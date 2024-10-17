package main

import "fmt"

//空接口作为函数的参数

func show(a interface{}) {
	fmt.Printf("值：%v 类型: %T \n", a, a)
}

func main(){
	show(20)
	show("你好golang")
	slice := []int{1,2,3,4}
	show(slice)


	//空接口作为切片
	var slice2 = []interface{}{"张三",20,true,32.2}
	fmt.Println(slice2)

	//空接口作为map值
	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "张三"
	studentInfo["age"] = "18"
	studentInfo["married"] = false
	fmt.Println(studentInfo)
}



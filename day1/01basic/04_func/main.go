package main

import "fmt"

func main () {


	fn := AddInt()
	fmt.Println(fn(1))
	fmt.Println(fn(1))
	fmt.Println(fn(1))
	fmt.Println(fn(1))

}


//函数语法： func 函数名 （接收参数   ）
func intSum(x, y int) int {
	return x + y
}

//闭包
//func（i int） int是个返回值，闭包的返回值是匿名函数
//闭包中的变量在函数调用完毕后不会销毁
func AddInt() func(i int) int {
	var num int
	return func(i int) int {
		num += 1
		return num
	}
}
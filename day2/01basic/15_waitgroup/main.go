package main

import (
	"fmt"
	"sync"
)

//第一步：定义一个计数器

var wg sync.WaitGroup

func main() {
	//100个协程的例子
	//wg.Add(100)
	//可以配合defer这样去使用
	//defer wg.Done()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go test1()
	}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go test2()
	}
	//没有先后顺序
	wg.Wait()
	fmt.Println("主进程退出")
}

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1() 你好golang")
	}
	wg.Done()
}



func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2() 你好golang")
	}
	wg.Done()
}
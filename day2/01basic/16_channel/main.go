package main

import "fmt"

func main()  {
	//channel的定义
	//ch1 := make(chan int, 10)
	//ch2 := make(chan bool, 4)
	//ch3 := make(chan []int, 3)

	ch := make(chan int,5)
	//向channel写入数据
	ch <- 10
	ch <- 12
	v1 := <- ch
	fmt.Println("v1", v1)
	//关闭channel
	close(ch)
	v2 := <- ch
	fmt.Println("v2", v2)
	//当channel关闭且没有值时，读出来的都是零值
	v3 := <- ch
	fmt.Println("v3", v3)
	//优雅的读channel，ch中没有值时会自动退出for循环
	//读channel
    for val := range ch{
    	fmt.Println(val)
	}
}



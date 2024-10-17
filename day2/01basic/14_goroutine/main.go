package main

import (
	"fmt"
	"time"
)

//使用goroutine
//主线程执行完，goroutine 自动退出
func main()  {
	go test()
	for i := 0; i < 2; i ++ {
        fmt.Println("main() 你好golang")
        time.Sleep(time.Microsecond * 100)
	}

}

func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好golang")
		time.Sleep(time.Microsecond * 100)
	}
}
package main

import "fmt"

//writeData
func writeDate(intChan chan int)  {
	for i :=0; i < 50; i++ {
		//写入数据
		intChan <- i
		fmt.Println("writeData ", i)
	}
	//当知道什么时候写完没能关闭尽量都关闭管道
	close(intChan)
}


//readData
func readData(intChan chan int, exitChan chan bool)  {
	for {
		//当读到零值时，ok=false
		v, ok := <- intChan
		if !ok {
			break
		}
		fmt.Println("readData ", v)
 	}
 	//读完后，写数据到exitChan中，相当于发生关闭信号给main主线程
 	exitChan <- true
 	close(exitChan)
}



func main()  {
   //定义intChan 和 exitChan
   //特性： 当一个channel同时被读写，它的大小不需要被设置到50,可以设置小一点（当前案例）
   intChan := make(chan int, 10)
   exitChan := make(chan bool, 1)
   go writeDate(intChan)
   go readData(intChan,exitChan)
   //监听退出信号并阻塞主线程，就是exitChan

   for {
   	_, ok := <- exitChan
   	if !ok {
   		break
	}
   }
   fmt.Println("主线程退出")
}

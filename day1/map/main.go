package main

import "fmt"

func main () {
	userInfo := map[string]string{
		"username": "root",
		"password": "123456",
	}

	fmt.Println(userInfo)


    userInfo2 := map[string]string{
    	"username": "zhangsan",
    	"password": "123456",
	}

	v,ok := userInfo2["username"]
	if ok {
		fmt.Println(v)

	} else {
		fmt.Println("map中没有此元素")
	}

    delete(userInfo2,"password")
	fmt.Println(userInfo2)

	soureMap := map[string]int{
		"zhangsan": 24,
		"lisi": 26,
		"wangwu": 24,
	}
    for k, v := range soureMap{
    	fmt.Println(k,v)
	}
	var a = 10

	fmt.Printf("%d \n", &a)   // 打印a的内存地址
	fmt.Printf("%d \n", *&a)  // 通过指针解引用，打印a的值
	fmt.Printf("%T \n", &a)   // 打印a的指针类型
	a11 := make([]int, 3, 10)              // 创建一个长度为3，容量为10的切片
	b := make(map[string]string)         // 创建一个空的map
	c := make(chan int, 1)               // 创建一个容量为1的channel

	fmt.Println(a11, b, c)                 // 输出a、b、c的值
}
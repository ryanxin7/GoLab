package main

import (
	"fmt"
)

func main () {
	//break
	k := 1
	for {
		if k <= 10 {
			fmt.Print(k)
		} else {
			fmt.Println("跳出循环")
			break
		}
		k++
	}

	//continue
	//当i=2的时候,continue 直接跳出当前循环，进行到i=3的循环
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	//return 跳出当前函数
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		if i == 5 {
			return
		}
	}
	fmt.Println("最后执行的打印")
}
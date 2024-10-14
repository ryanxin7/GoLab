package main

import "fmt"

func main ()  {
	//第一种
	for i :=0; i <10; i++ {
		fmt.Println(i)
	}
	//第二种
	i := 10
	for (i < 10) {
		fmt.Println(i)
		i++
	}

	//无限for循环
//	for {
//		fmt.Println(1)
//	}

	//for range 循环
	str := "abc上海"
	for i, v := range str {
		fmt.Printf("索引=%d,值=%c \n", i, v)
	}

	//switch case
	score := "D"
	switch score {
	case "A":
		fmt.Println("非常棒")
	case "B":
		fmt.Println("优秀")
	case "C":
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}
}
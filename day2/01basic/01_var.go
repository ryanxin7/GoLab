package main

import "fmt"

func main() {

	fmt.Print("zhangsan", "lisi", "wangwu")   // 不会有空格，不会换行
	fmt.Println("zhangsan", "lisi", "wangwu") // 自动有空格并换行

	name := "zhangsan"
	age := 20
	fmt.Printf("%s 今年 %d 岁", name, age) // 格式化输出

	info := fmt.Sprintf("姓名: %s, 性别: %d", name, 20) // 格式化并返回字符串
	fmt.Println(info)                               // 输出info字符串

}

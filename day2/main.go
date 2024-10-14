package main

import (
	encapsulation "example.com/day2/01basic/09_encapsulation"
	"fmt"
)

func main() {
	s := encapsulation.NewStudent("张三", 80)
	fmt.Println(s)
	fmt.Println(s.GetScore())
	s.SetScore(90)
	fmt.Println(s.GetScore())
}

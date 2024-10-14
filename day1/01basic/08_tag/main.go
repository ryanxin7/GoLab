package main

import (
	"encoding/json"
	"fmt"
)
//使用场景：使用gin框架通过api接口返回出去的结构体，要使用json，因为gin会做序列化处理
//tag中json用于序列化和反序列化
type Student struct {
	ID int `json:"id"`
	Gender string `json:"gender"`
	name string //私有属性不能被json包访问
	Sno string `json:"sno"`
}
func main()  {
	var s1 = &Student{
		ID:     1,
		Gender: "男",
		name:   "张三",
		Sno:    "s0001",
	}
	fmt.Println(s1)
	//序列化
	strByte, err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(strByte))

	//反序列化
	str := `{"id":3,"gender":"女","sno":"s0003"}`
	var s2 = &Student{}
	err = json.Unmarshal([]byte(str), s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s2)

}
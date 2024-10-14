package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID int
	Gender string
	Name string  //私有属性不能被json包访问
	Sno string
}

func main()  {
	var s1 = &Student{
		ID:     1,
		Gender: "男",
		Name:   "张三",
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
	str := `{"ID":1,"Gender":"男","Name":"张三","Sno":"s0001"}`
	var s2 = &Student{}
	err = json.Unmarshal([]byte(str), s2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s2)

}
package main

import (
	"fmt"
	"encoding/json"
)

type Monster struct {
	Name string `json:"name"` //TODO 使用到反射-后面详解
	Age int `json:"age"`
	Skill string `json:"skill"`
}

func main()  {
	//TODO struct 的每个字段上，可以写一个 tag， 该tag 可以通过反射机制获取
	//TODO 常见的使用场景就是序列化和反序列化
	monster := Monster{"牛魔王", 500, "芭蕉扇"}

	//monster 序列化成 json 字符串
	jsonMonster, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 错误")
	}
	fmt.Println("jsonMonster=", jsonMonster)
	fmt.Println("jsonMonster=", string(jsonMonster))

}
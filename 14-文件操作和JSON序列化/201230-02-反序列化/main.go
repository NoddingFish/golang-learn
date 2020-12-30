package main

import (
	"fmt"
	"encoding/json"
)

//TODO 重要：
//TODO 1、要确保反序列化后的数据类型和原来序列化前的数据类型一致
//TODO 2、如果 json 字符串不需要转义 \"

/*
jsonStruct 序列化后： {"name":"牛魔王","monster_age":500,"Birthday":"901-1-1","Sal":20000,"Skill":"牛魔拳"}
jsonMap 序列化后： {"address":"北京","age":100,"name":"红孩儿"}
jsonSlice 序列化后： [{"address":"北京","age":10,"name":"jack"},{"address":["上海","夏威夷"],"age":20,"name":"mary"}]
*/

type Monster struct {
	Name string `json:"name"`
	Age int `json:"monster_age"`
	Birthday string
	Sal float64
	Skill string
}

func unmarshalStruct()  {
	str := "{\"name\":\"牛魔王\",\"monster_age\":500,\"Birthday\":\"901-1-1\",\"Sal\":20000,\"Skill\":\"牛魔拳\"}"

	var monster Monster

	//序列化 func Marshal(v interface{}) ([]byte, error)
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("unmarshalStruct 序列化失败：", err)
	}
	fmt.Println("unmarshalStruct 反序列化后：", monster)
	fmt.Println("monster.Name=", monster.Name)
}


func unmarshalMap()  {
	str := "{\"address\":\"北京\",\"age\":100,\"name\":\"红孩儿\"}"

	var map1 map[string]interface{}


	//反序列化
	//注意：反序列化不需要make，make被封装到 Unmarshal 函数中
	err := json.Unmarshal([]byte(str), &map1)
	if err != nil {
		fmt.Println("unmarshalMap 序列化失败：", err)
	}
	fmt.Println("unmarshalMap 反序列化后：", map1)
}

func unmarshalSlice()  {
	str := "[{\"address\":\"北京\",\"age\":10,\"name\":\"jack\"}," + 
	"{\"address\":[\"上海\",\"夏威夷\"],\"age\":20,\"name\":\"mary\"}]"

	var slice []map[string]interface{}


	//反序列化
	//TODO 注意：反序列化不需要make，make被封装到 Unmarshal 函数中
	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("unmarshalSlice 序列化失败：", err)
	}
	fmt.Println("unmarshalSlice 反序列化后：", slice)
}

func main()  {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}

package main

import (
	"fmt"
	"encoding/json"
)

func main()  {
	jsonStruct()
	jsonMap()
	jsonSlice()
	jsonFloat64() //TODO 基本数据类型序列化，没有区别，意义不大
}

type Monster struct {
	// name string //! 不能直接小写， json 包不能访问到这个字段，序列化会丢失这个字段
	Name string `json:"name"` //TODO 定义 tag  ==> 反射机制
	Age int `json:"monster_age"`
	Birthday string
	Sal float64
	Skill string
}

func jsonStruct()  {
	monster := Monster{
		Name: "牛魔王",
		Age: 500,
		Birthday: "901-1-1",
		Sal: 20000.0,
		Skill: "牛魔拳",
	}
	//TODO 注意： Name 应该是大写，但是可能需求是小写 ==> 定义的时候需要 tag
	//! {"Name":"牛魔王","Age":500,"Birthday":"901-1-1","Sal":20000,"Skill":"牛魔拳"}
	//! {"name":"牛魔王","monster_age":500,"Birthday":"901-1-1","Sal":20000,"Skill":"牛魔拳"}

	//序列化 func Marshal(v interface{}) ([]byte, error)
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("jsonStruct 序列化失败：", err)
	}
	fmt.Println("jsonStruct 序列化后：", string(data))
}

func jsonMap()  {
	var map1 map[string]interface{}
	map1 = make(map[string]interface{})

	map1["name"] = "红孩儿"
	map1["age"] = 100
	map1["address"] = "北京"


	//序列化 func Marshal(v interface{}) ([]byte, error)
	data, err := json.Marshal(&map1)
	if err != nil {
		fmt.Println("jsonMap 序列化失败：", err)
	}
	fmt.Println("jsonMap 序列化后：", string(data))
}

func jsonSlice()  {
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 10
	m1["address"] = "北京"
	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})
	m2["name"] = "mary"
	m2["age"] = 20
	m2["address"] = []string{"上海", "夏威夷"}
	slice = append(slice, m2)

	//序列化 func Marshal(v interface{}) ([]byte, error)
	data, err := json.Marshal(&slice)
	if err != nil {
		fmt.Println("jsonSlice 序列化失败：", err)
	}
	fmt.Println("jsonSlice 序列化后：", string(data))
}

func jsonFloat64() {
	var num1 float64 = 234.56

	//序列化 func Marshal(v interface{}) ([]byte, error)
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("jsonFloat64 序列化失败：", err)
	}
	fmt.Println("jsonFloat64 序列化后：", string(data))
}


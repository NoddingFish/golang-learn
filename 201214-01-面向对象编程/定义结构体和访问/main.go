package main

import (
	"fmt"
)

type Monster struct {
	Name string
	Age int
}

func main()  {
	fmt.Println("========================== 定义结构体2 ==========================")
	// mon1 := Monster{}
	mon1 := Monster{"牛魔王", 500}
	mon1.Age = 501
	fmt.Println("mon1=", mon1)

	fmt.Println("========================== 定义结构体3 - 结构体指针 ==========================")
	// mon1 := Monster{}
	mon3 := new(Monster)
	//以下是标准写法
	(*mon3).Name = "白骨精"
	(*mon3).Age = 501
	fmt.Println("标准写法 mon3=", *mon3)

	//简化写法 - 底层做了处理
	mon3.Name = "白骨精2"
	mon3.Age = 502
	fmt.Println("简化写法 mon3=", *mon3)

	fmt.Println("========================== 定义结构体4 - 结构体指针 ==========================")
	// var mon4 *Monster = &Monster{"银角大王", 180}
	mon4 := &Monster{"银角大王", 180}
	fmt.Println("mon4=", *mon4)
	//以下是标准写法
	(*mon4).Name = "金角大王"//TODO 不能 *mon4.Name 这样写，因为 . 的优先级比 * 高
	(*mon4).Age = 201
	fmt.Println("标准写法 mon4=", *mon4)

	//简化写法 - 底层做了处理
	mon4.Name = "金角大王2"
	mon4.Age = 202
	fmt.Println("简化写法 mon4=", *mon4)

}
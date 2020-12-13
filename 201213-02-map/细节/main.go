package main

import (
	"fmt"
)

func modi(map1 map[int]int)  {
	map1[10] = 1000
}

type Stu struct {
	Name string
	Age int
	Address string
}

func main()  {
	fmt.Println("============================= map细节 =============================")
	map1 := make(map[int]int, 1)
	map1[10] = 10
	map1[1] = 100
	map1[8] = 800
	map1[4] = 400

	fmt.Println("map1=", map1)

	//TODO 1、引用类型
	modi(map1)
	fmt.Println("map1=", map1)//说明map是引用类型

	//TODO 2、使用不需要扩容，可以自动扩容
	//TODO 3、map 的 value 经常使用 struct 结构体
	student := make(map[string]Stu, 0)
	stu1 := Stu{"tom", 18, "北京市"}
	stu2 := Stu{"mary", 22, "上海市"}
	student["n1"] = stu1
	student["n2"] = stu2
	fmt.Println(student)

	for index, value := range student {
		fmt.Print("学生", index, "的姓名是：", value.Name, "\n")
		fmt.Print("学生", index, "的年龄是：", value.Age, "\n")
		fmt.Print("学生", index, "的地址是：", value.Address, "\n")
		fmt.Println()
	}

}
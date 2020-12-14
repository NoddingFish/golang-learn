package main

import (
	"fmt"
)

//TODO 结构体的声明 结构体首字母大写，说明可以在包外被使用，下相同
type Cat struct {
	Name string // TODO 字段的类型可以试基本类型、数组或引用类型
	Age int
	Color string
	isFish bool
	scores [2]int
	ptr1 *int
	slice1 []float64
	map1 map[string]string
}

type Monster struct {
	Name string
	Age int
}

func main()  {
	//TODO 结构体struct是值类型
	var cat01 Cat
	fmt.Println("cat01=", cat01)//不赋值，会有默认值
	cat01.Name = "小白"
	cat01.Age = 30
	cat01.Color = "白色"
	fmt.Println("cat01=", cat01)
	fmt.Println("cat01.Name=", cat01.Name)
	fmt.Println("cat01.Age=", cat01.Age)
	fmt.Println("cat01.Color=", cat01.Color)

	//TODO 指针 slice 和 map 的零值都是nil，即还没有分配空间，需要make
	if cat01.slice1 == nil {
		fmt.Println("cat01.slice01=", "nil")
	}
	cat01.slice1 = make([]float64, 2)
	cat01.slice1[0] = 1.2//会报错，需要先make
	fmt.Println("cat01=", cat01)


	fmt.Println("================================== 值类型 ==================================")
	var monster1 Monster
	monster1.Name = "牛魔王"
	monster1.Age = 500

	monster2 := monster1
	monster2.Name = "白骨精"
	monster2.Age = 200

	fmt.Println("monster1=", monster1)
	fmt.Println("monster2=", monster2)	

}
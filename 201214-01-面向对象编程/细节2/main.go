package main

import (
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int // 可以转换
	// Num1 int // 不能转换
	// name int // 不能转换
	// Num float64 // 不能转换
}

func main()  {
	//TODO 2、结构体是用户单独定义的类型，和其他的类型转换时需要有完全相同的字段（名字、个数和类型）
	fmt.Println("========================== 结构体是用户单独定义的类型，和其他的类型转换时需要有完全相同的字段（名字、个数和类型）")
	var a A
	var b B
	a = A(b)
	fmt.Println("a=", a)

	//TODO! 2、结构体进行 type 重新订单（相当于取别名）,go认为是新的数据类型，但是可以相互间可以强制转换
}
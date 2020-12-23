package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main()  {
	//TODO 类型断言：由于接口是一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言
	var a interface{}
	var point Point = Point{1, 2}
	a = point //ok
	//如何将 a 赋给一个 Point 变量？
	var b Point
	// b = a//!报错
	b = a.(Point) //TODO 类型断言
	//TODO 表示判断a是否是指向Point类型的变量，如果是就转成Point类型并赋给b变量，否则报错
	fmt.Println("b=", b)


	var x interface{}
	var b2 float32 = 1.1
	x = b2
	y := x.(float32)
	// y := x.(float64)//报错
	fmt.Printf("y 的类型是%T 值是%v \n", y, y)


	//如何在进行断言时，带上检测机制,如果成功就ok，否则也不要报 panic
	fmt.Println("=================== 带检测的类型断言 ===================")
	// y2, ok := x.(float64)
	// if ok {
	if y2, ok := x.(float64);ok {
		fmt.Println("convert success")
		fmt.Printf("y2 的类型是%T 值是%v", y2, y2)
	} else {
		fmt.Println("convert fail")
	}
	// y := x.(float64)//报错
	
}
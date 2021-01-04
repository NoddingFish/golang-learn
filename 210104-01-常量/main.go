package main

import (
	"fmt"
)

func main()  {
	var num int
	fmt.Println("num=", num)

	num = 9

	//TODO 1、常量定义的时候，必须赋值
	//TODO 2、常量不能修改
	//TODO 3、常量只能是 bool、数值类型、string类型
	//TODO 4、首字母大小写，控制使用的范围
	const tax int = 0
	fmt.Println("tax=", tax)

	const tax2 int = 9 / 3
	// const tax2 int = num / 3//!报错，num是 变量，认为是可变的
	fmt.Println("tax2=", tax2)


	//TODO 专业写法
	const (
		a = iota //默认为0
		b 	     //b在a基础上加1
		c 	     //c在b基础上加1
		d 	     //d在c基础上加1
	)
	fmt.Println("a=", a)
	fmt.Println("b=", b)
	fmt.Println("c=", c)
	fmt.Println("d=", d)

	const (
		a2 = iota
		b2 = iota
		c2,d2 = iota, iota
	)
	fmt.Println("a2=", a2)
	fmt.Println("b2=", b2)
	fmt.Println("c2=", c2)
	fmt.Println("d2=", d2)
}
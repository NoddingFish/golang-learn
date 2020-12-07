package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//TODO　内置函数

	//TODO 1、len 求长度，string array slice map channel
	fmt.Println("len=", len("1329背景"))

	//TODO 2、new 分配内存, 主要用来分配值类型, int, float, struct...
	num1 := 100
	fmt.Printf("num1的类型是：%T , 值是：%v, 地址是：%v \n", num1, num1, &num1)
	num2 := new(int)
	fmt.Printf("num2的类型是：%T , 值是：%v, 地址是：%v, num2指针的值是：%v \n", num2, num2, &num2, *num2)
	*num2 = 100
	fmt.Printf("num2的类型是：%T , 值是：%v, 地址是：%v, num2指针的值是：%v \n", num2, num2, &num2, *num2)

	//TODO 3、make 分配内存, 主要用来分配引用类型, channel, map, slice - 后面讲解
}
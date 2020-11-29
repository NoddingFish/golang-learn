package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//TODO 赋值运算符
	fmt.Println("======================= 赋值运算符 =====================")
	var i int
	i = 10 // 赋值
	fmt.Printf("i = %v \n", i)

	// a b 值交换
	a := 9
	b := 2
	t := a
	a = b
	b = t
	fmt.Printf("交换后的情况是 a = %v , b = %v \n", a, b)

	i += 9
	fmt.Printf("i += 9 = %v \n", i)

	//练习：
	//a1 b1 值交换 , 不能使用中间变量
	var a1 int = 100
	var b1 int = 200
	a1 = a1 + b1
	b1 = a1 - b1
	a1 = a1 - b1
	fmt.Printf("交换后的情况是 a1 = %v , b1 = %v \n", a1, b1)
}
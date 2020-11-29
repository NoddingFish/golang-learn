package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//TODO 其他运算符
	fmt.Println("======================= 其他运算符 =====================")
	var i int
	var pre *int
	i = 10 // 赋值

	pre = &i

	fmt.Printf("i的地址 = %v \n", &i)
	fmt.Printf("pre = %v \n", *pre)

	//练习：
	//两个数的最大值 和 三个数的最大值
	var num1 int = 29
	var num2 int = 9
	var num3 int = 80
	var max int
	if num1 > num2 {
		max = num1
	} else {
		max = num2
	}
	fmt.Println("两个数的最大值是：", max)

	if max < num3 {
		max = num3
	}
	fmt.Println("三个数的最大值是：", max)
}
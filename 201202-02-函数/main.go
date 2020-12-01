package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//TODO 函数
	var num1 float64 = 2.6
	var num2 float64 = 2.2
	res := suanshu(num1, num2, '+')
	fmt.Println("结果是：", res)
}

func suanshu(num1 float64, num2 float64, opa byte) float64 {
	var res float64
	switch opa {
		case '+':
			res = num1 + num2
		case '-':
			res = num1 - num2
		case '*':
			res = num1 * num2
		case '/':
			res = num1 / num2
		default:
			fmt.Println("输入有误")
	}
	return res
}
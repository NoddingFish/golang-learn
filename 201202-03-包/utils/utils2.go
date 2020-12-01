package utils

import "fmt"

var Num2 int = 98882

func Suanshu2(num1 float64, num2 float64, opa byte) float64 {
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
			fmt.Println("输入错误")
	}
	return res
}
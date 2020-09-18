package main

import (
	"fmt"
)

func main() {
	//浮点型都是有符号的

	//浮点型都可能出现精度损失的
	
	var price float32 = -8922.12123531123131111112138

	var price2 float64 = -8922.12123531123134532138

	fmt.Println("price=", price)
	fmt.Println("price2=", price2)

	//Golang 的浮点型默认是 float64
	var price3 = 1.2

	fmt.Printf("Golang 的浮点型默认：%T \n", price3)


	num1 := 1.1
	num2 := .221 // .221 => 0.221

	fmt.Println("num1=", num1, "num2=", num2)

	num3 := 2.21454e3 // 2.21 * 10的3次方
	num4 := 2.21454E3 // 2.21 * 10的3次方
	num5 := 2.21454E-1 // 2.21 * 10的3次方

	fmt.Println("num3=", num3)
	fmt.Println("num4=", num4)
	fmt.Println("num5=", num5)

}
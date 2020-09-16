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


}
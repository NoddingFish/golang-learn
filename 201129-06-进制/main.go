package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//二进制，在 golang 中，不能直接使用二进制
	var num1 int = 5
	fmt.Printf("%v 的二进制是 %b \n", num1, num1)

	var num2 int = 011
	fmt.Println("八进制 num2=", num2)

	var num3 int = 0x11
	fmt.Println("十六进制 num3=", num3)

	// fmt.Println("=================== 其他进制转十进制 ===================")

	// 二进制转八进制和十六进制
	//11 100 101 转八进制： 0 3 4 5 = 0345
	//11 1001 0110 转十六进制：ox 3 9 6 = ox396

	// 八进制和十六进制转二进制
	// 八进制 075 => 二进制： 111 101
	// 十六进制 0xAF => 二进制： 1010 1111
}
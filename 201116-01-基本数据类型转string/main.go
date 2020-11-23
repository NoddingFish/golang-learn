package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
	"strconv"
)

func main() {
	//第一种方式 fmt.Sprintf()
	var num1 int64 = 99
	var num2 float32 = 23.892
	var b bool
	var chat byte = 'j'

	var str string

	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str type %T  str=%q \n", str, str)

	str = fmt.Sprintf("%f", num2)
	fmt.Printf("str type %T  str=%q \n", str, str)

	str = fmt.Sprintf("%t", b)
	fmt.Printf("str type %T  str=%q \n", str, str)

	str = fmt.Sprintf("%c", chat)
	fmt.Printf("str type %T  str=%q \n", str, str)

	//第一种方式 strconv 函数
	var num3 int32 = 99
	var num4 float64 = 23.892
	var b2 bool = true
	
	str = strconv.FormatInt(int64(num3), 10)
	fmt.Printf("str type %T  str=%q \n", str, str)

	str = strconv.FormatFloat(num4, 'f', 10, 64)
	fmt.Printf("str type %T  str=%q \n", str, str)

	str = strconv.FormatBool(b2)
	fmt.Printf("str type %T  str=%q \n", str, str)

	var num5 int32 = 78
	str = strconv.Itoa(int(num5))
	fmt.Printf("str type %T  str=%q \n", str, str)

}
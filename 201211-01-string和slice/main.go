package main

import (
	"fmt"
)

func main()  {
	//TODO string底层是一个byte数组，因此string可以进行切片处理
	str := "hello@Ran世界"
	//使用切片获取到 Ran
	slice01 := str[6:]
	fmt.Println("str    =", str)
	fmt.Println("slice01=", slice01)

	//TODO string本身是不可变的，不能通过 str[0] = 'z' 修改 
	// str[0] = 'z'//! cannot assign to str[0]

	//TODO string 修改 string->[]byte / []rune -> 修改 -> 转成 string
	arr01 := []byte(str)
	arr01[0] = 'z'
	str = string(arr01)
	fmt.Println("str    =", str)
	//* 上面可以处理英文和数字-不能处理中文
	// arr01[0] = '中'//! 报错

	arr02 := []rune(str)
	arr02[0] = '中'
	str = string(arr02)
	fmt.Println("str    =", str)

}
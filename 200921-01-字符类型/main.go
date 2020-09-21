package main

import (
	"fmt"
)

func main() {
	//Golang 中没有专门的字符类型。如果存储单个的字符，使用 byte 类型
	//Go的字符串不同，它是由字节组成的。

	var str1 byte = 'a'
	var str2 byte = '0' //字符的0

	fmt.Println("str1=", str1)
	fmt.Println("str2=", str2)

	//如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("str1=%c str2=%c \n", str1, str2)

	// var str3 byte = '北'//overflow
	var str3 int = '北'//overflow
	fmt.Printf("str3=%c\n", str3)


	// var str3 = "i am chinese";

	// fmt.Printf("字符类型的默认类型：%T \n", str3)

}
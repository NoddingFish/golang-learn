package main

import (
	"fmt"
)

func main() {
	// string 的基本使用
	var str1 string = "背景颜色"

	fmt.Println(str1)

	str2 := "hello"
	//str2[0] = "i"

	fmt.Println(str2)

	str3 := "abc\nabc"
	fmt.Println(str3)


	//使用反引号 ``
	str4 := `package main

	import (
		"fmt"
	)
	
	func main() {
		// string 的基本使用
		var str1 string = "背景颜色"
	
		fmt.Println(str1)
	
		str2 := "hello"
		//str2[0] = "i"
	
		fmt.Println(str2)
	
		//使用反引号
		
	}`

	fmt.Println(str4)

	str51 := "hello "
	var str52 string = "world"
	fmt.Println(str51 + str52) 

	//当拼接很长的时候，分行写，但是 + 每行必须写在最后
	str6 := "hello " + " world1" + " world1" + " world1" + 
	" world1" + " world1" + " world1" + " world1"
	fmt.Println(str6) 
}
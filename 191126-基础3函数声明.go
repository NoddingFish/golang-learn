package main

import (
	"fmt"
	"os"
)

func main()  {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	fmt.Println("it's over 9000!", os.Args[1])// 索引 0 处 -- 总是当前可运行程序的路径。
	a, b := test(1, 2)

	fmt.Println(a)
	fmt.Println(b)

	//go run 191126-2导入包.go  99123
}


func test(a, b int) (int, bool) {
	var isCheck bool = true
	c := a + b
	if c < 0 {
		isCheck = false
	}
	return c, isCheck
}
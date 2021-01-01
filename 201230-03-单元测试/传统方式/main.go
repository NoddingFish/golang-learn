package main

import (
	"fmt"
)

func addUpper(n int) int  {
	res := 0
	for i := 1; i <= n - 1; i++ {
		res += i
	}
	return res
}

func main() {
	res := addUpper(10)

	if res != 55 {
		fmt.Println("addUpper错误，返回值是：", res, " 期望值是： 55")
	} else {
		fmt.Println("addUpper正确，返回值是：", res)
	}
}
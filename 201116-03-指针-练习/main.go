package main

import (
	"fmt"
)

func main() {
	var num int = 89
	fmt.Printf("num的地址是：%v\n", &num)

	var ptr *int = &num
	*ptr = 99
	fmt.Printf("num的修改后的值是：%v\n", num)
}
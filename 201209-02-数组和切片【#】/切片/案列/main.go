package main

import (
	"fmt"
)

func test(slice []int) {
	slice[0] = 100
}

func main()  {
	fmt.Println("============================ 切片 引用类型 ============================")
	var slice01 []int
	var arr = [...]int{1,2,3,4,5}
	slice01 = arr[:]
	var slice02 = slice01
	slice02[0] = 10
	fmt.Println("slice01=", slice01)
	fmt.Println("slice02=", slice02)
	fmt.Println("arr    =", arr)

	var slice03 = []int{1,2,3,4,5}
	fmt.Println("slice03=", slice03)
	test(slice03)
	fmt.Println("slice03=", slice03)
}
package main

import (
	"fmt"
)

func main()  {
	fmt.Println("============================== 练习1 ==============================")
	var byte1 [26]byte
	for i := 0; i < 26; i++ {
		byte1[i] = 'A' + byte(i)
	}
	for i := 0; i < len(byte1); i++ {
		fmt.Printf("%c ", byte1[i])
	}
	fmt.Println()
	fmt.Println()
	
	fmt.Println("============================== 练习2 ==============================")
	//请求一个数组的最大值，并得到对应的下标
	arr02 := [...]int{20, -2, 18, 9, 90, 810, 2, 11}
	maxVal := arr02[0]
	maxIndex := 0
	for index, value := range arr02 {
		if index > 0 && maxVal < value {
			maxVal = value
			maxIndex = index
		}
	}
	fmt.Printf("数组的最大值是：%v 对应的下标是：%v \n", maxVal, maxIndex)
}
package main

import (
	"fmt"
	"math/rand"
	"time"
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
	fmt.Println()

	fmt.Println("============================== 练习3 ==============================")
	arr03 := [...]float64{128, 169, 220, 156, 199.7}
	var sum03 float64
	var avg03 float64
	for _, value := range arr03 {
		sum03 += value
	}
	avg03 = sum03 / float64(len(arr03))
	fmt.Printf("arr03的和=%v 平均值=%.2f \n", sum03, avg03)
	fmt.Println()

	fmt.Println("============================== 练习4 - 数组反转 ==============================")

	var arr04 [5]int
	len := len(arr04)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len; i++ {
		arr04[i] = rand.Intn(100)
	}
	fmt.Println("交换前~：", arr04)
	temp := 0//临时变量，用于交换
	for i := 0; i < len / 2; i++ {
		temp = arr04[len - 1 - i]
		arr04[len - 1 - i] = arr04[i]
		arr04[i] = temp
	}
	fmt.Println("交换后~：", arr04)
}
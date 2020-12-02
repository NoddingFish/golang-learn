package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//函数调用的机制
	var num1 float64 = 85
	fmt.Println(jia(num1))
	fmt.Println(num1)

	fmt.Println("====================== 函数多返回值 ======================")
	// res1, _ := sShu(10, 8) // 忽略一个值，可以使用 _ 占位
	res1, res2 := sShu(10.5, 8.2)
	fmt.Printf("两数之和：%v 两数之差是：%.2f \n", res1, res2)

	fmt.Println("====================== 函数递归调用1 ======================")
	diGui(4)

	fmt.Println("====================== 函数递归调用2 ======================")
	diGui2(4)

}

func jia(num1 float64) float64 {
	return num1 + 1
}

func sShu(num1 float64, num2 float64) (float64, float64) {
	res1 := num1 + num2
	res2 := num1 - num2
	return res1, res2
}

func diGui(num int)  {
	if num > 2 {
		num--
		diGui(num)
	}
	fmt.Println("num1=", num)
}

func diGui2(num int)  {
	if num > 2 {
		num--
		diGui2(num)
	} else {
		fmt.Println("num2=", num)
	}
}
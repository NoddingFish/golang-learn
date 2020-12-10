package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)


func main() {
	num1 := 10.7333333
	num2 := fmt.Sprintf("%.2f", num1)//保留小数点后两位有效，四舍五入
	fmt.Println("num2=", num2)

	//TODO 1、数组 是值类型
	fmt.Println("============================ 数组 ============================")
	var arr1 [6]float64
	fmt.Println("数组的初始值：", arr1)
	fmt.Printf("数组arr的地址：%p \n", &arr1)
	fmt.Printf("数组arr[0]的地址：%p \n", &arr1[0])
	fmt.Printf("数组arr[1]的地址：%p \n", &arr1[1])
	fmt.Printf("数组arr[2]的地址：%p \n", &arr1[2])
	fmt.Printf("数组arr[3]的地址：%p \n", &arr1[3])
	fmt.Printf("数组arr[4]的地址：%p \n", &arr1[4])
	//TODO 数组的各个元素的地址间隔是依据数组的类型决定的，比如int => 8字节 int32 => 4字节
	arr1[0] = 3.0
	arr1[1] = 5.0
	arr1[2] = 1.0
	arr1[3] = 3.4
	arr1[4] = 2.0
	arr1[5] = 50.0
	totalWeight := 0.0
	for i := 0; i < len(arr1); i++ {
		totalWeight += arr1[i]
	}
	// avgWeight := totalWeight / 6  //这里写 6 是可以的，因为变量有类型，6是个具体的数，没有类型
	avgWeight := totalWeight / float64(len(arr1))
	fmt.Printf("总体重=%v 平均体重=%.2f \n", totalWeight, avgWeight)


	//练习
	// var score [5]float64
	// for i := 0; i < len(score); i++ {
	// 	fmt.Printf("请输入第%d个元素的值：\n", i + 1)
	// 	fmt.Scanln(&score[i])
	// }
	// fmt.Println("数组是：", score)


	fmt.Println("======================== 初始化数组 ========================")
	var initArr1 [3]int = [3]int{1, 2, 3}
	fmt.Println("initArr1 = ", initArr1)

	var initArr2 = [3]int{1, 2, 3}
	fmt.Println("initArr2 = ", initArr2)

	var initArr3 = [...]float64{1, 2, 3, 2.3}
	fmt.Println("initArr3 = ", initArr3)

	var initArr4 = [...]float64{1: 1.9, 2: 99, 0: 900}//指定下标
	fmt.Println("initArr4 = ", initArr4)

	initArr5 := [...]string{1: "tom", 2: "tes", 0: "john"}
	fmt.Println("initArr5 = ", initArr5)

	fmt.Println("======================== 数组遍历 ========================")
	//1、for 遍历，常规的方式
	//2、for-range 遍历 不需要的值可以使用 _ 占位符
	for index1, value1 := range initArr5 {
		fmt.Println("initArr5 = ", index1, value1)
	}


}
package main

import (
	"fmt"
)

func main() {
	//TODO 二维数组
	fmt.Println("============================ 二维数组 ============================")
	var arr01 [4][6]int
	fmt.Println(arr01)
	arr01[1][2] = 1
	arr01[2][1] = 2
	arr01[2][3] = 3
	for _, value := range arr01 {
		for _, value2 := range value {
			fmt.Print(value2, " ")
		}
		fmt.Println()
	}

	fmt.Printf("arr01[0]的地址是： %p \n", &arr01[0])
	fmt.Printf("arr01[1]的地址是： %p \n", &arr01[1])

	fmt.Printf("arr01[0][0]的地址是： %p \n", &arr01[0][0])
	fmt.Printf("arr01[1][0]的地址是： %p \n", &arr01[1][0])


	fmt.Println("============================ 二维数组 - 使用2 ============================")
	// arr02 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// arr02 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	// var arr02 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var arr02 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr02=", arr02)

	// var num1 = 1
	var num1 float64 = 1
	fmt.Println("num1=", num1)

	fmt.Println("============================ 多维数组 ============================")
	var arr03 [4][6][2]int
	fmt.Println(arr03)

	var slice01 = make([][]int, 2)
	slice01[0] = []int{1, 2, 3}
	fmt.Println(slice01)

	fmt.Println("============================ 案列 ============================")
	var scores [3][5]float64
	for i, v := range scores {
		for j, _ := range v {
			fmt.Printf("请输入第%v个班级第%v个学生的成绩：", i + 1, j + 1)
			fmt.Scanln(&scores[i][j])
		}
		fmt.Println()
	}
	fmt.Println("scores=", scores)
	for i, v := range scores {
		for j, _ := range v {
			fmt.Printf("请输入第%v个班级第%v个学生的成绩：",
			 i + 1,
			 j + 1)
			fmt.Scanln(&scores[i][j])
		}
		fmt.Println()
	}
}
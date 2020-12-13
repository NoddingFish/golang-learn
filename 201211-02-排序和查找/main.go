package main

import (
	"fmt"
)

func twoSearch(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	middleIndex := (leftIndex + rightIndex) / 2
	middle := (*arr)[middleIndex]
	if (findVal < middle) {
		twoSearch(arr, leftIndex, middleIndex - 1, findVal)
	} else if (findVal == middle) {
		fmt.Println("找到了，下标为：", middleIndex)
	} else {
		twoSearch(arr, middleIndex + 1, rightIndex, findVal)
	}
}

func main() {
	//TODO 数组的排序
	//内部排序 交换式排序法 选择式排序法 插入式排序法
	fmt.Println("============================ 交换排序-冒泡排序 ============================")
	//冒泡排序：前面的数和后面的数比较，前面的数大，则交换
	//     [24, 69, 80, 57, 13]
	// ==== 第一轮
	//1 => [24, 69, 80, 57, 13]
	//2 => [24, 69, 80, 57, 13]
	//3 => [24, 69, 57, 80, 13]
	//4 => [24, 69, 57, 13, 80]
	// ==== 第二轮
	//1 => [24, 69, 57, 13, 80]
	//2 => [24, 57, 69, 13, 80]
	//3 => [24, 57, 13, 69, 80]
	// ==== 第三轮
	//1 => [24, 57, 13, 69, 80]
	//2 => [24, 13, 57, 69, 80]
	// ==== 第四轮
	//1 => [13, 24, 57, 69, 80]
	arr01 := [...]int{24, 69, 80, 57, 13}
	fmt.Println("排序前=", arr01)
	for ceng := 0; ceng < len(arr01) - 1; ceng++ {
		for i := 0; i < len(arr01) - ceng - 1; i++ {
			// fmt.Println("i=", i)
			if arr01[i] > arr01[i + 1] {
				temp := arr01[i + 1]
				arr01[i + 1] = arr01[i]
				arr01[i] = temp
			}
			fmt.Printf("第%v轮第%v次排序=%v \n", ceng+1, i+1, arr01)
		}
	}
	fmt.Println("排序后=", arr01)


	//外部排序 合并排序法和直接合并排序法


	//TODO 查找
	//1、顺序查找
	//2、二分查找
	fmt.Println("============================ 查找 - 二分查找 ============================")
	arr02 := [6]int{1, 8, 10, 89, 1000, 1234}
	
	fmt.Println("请输入要查找的数字：")
	num := 0
	fmt.Scanln(&num)
	twoSearch(&arr02, 0, len(arr02) - 1, num)
}
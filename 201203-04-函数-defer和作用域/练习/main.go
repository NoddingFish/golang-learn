package main

import (
	"fmt"
)

func Jzt(level int) bool {
	
	for i := 1; i <= level; i++ {
		for j := 0; j < level - i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2 * i - 1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	return true
}

func JiuJiu(num91 int) {
	for i := 1; i <= num91; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d = %d \t", j, i, j * i)
		}
		fmt.Println()
	}
}

func main()  {
	//打印空心金字塔
	var level int
	fmt.Printf("输入金字塔的层级：")
	fmt.Scanln(&level)
	Jzt(level)
	fmt.Printf("================= 九九乘法表 ================= \n")
	JiuJiu(9)
}
package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//多重循环控制 - 一般使用两层
	//练习一
	// allSum := 0.0
	// for i := 0; i < 3; i++ {
	// 	sum := 0.0
	// 	for j := 0; j < 5; j++ {
	// 		var score float64
	// 		fmt.Printf("请输入第%d个学生的成绩：", j+1)
	// 		fmt.Scanln(&score)
	// 		sum += score
	// 	}
	// 	fmt.Printf("班级 %d 的平均分是 %v \n", i, sum / 5)
	// 	allSum += sum
	// }
	// fmt.Printf("所有班级的总平均分是%v \n", allSum / 15)

	//练习二
	//打印金字塔
	var jinzita1 int
	fmt.Printf("输入金字塔（一）的层级：")
	fmt.Scanln(&jinzita1)
	for i := 1; i <= jinzita1; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//练习三
	//打印金字塔
	var jinzita2 int
	fmt.Printf("输入金字塔（二）的层级：")
	fmt.Scanln(&jinzita2)
	for i := 1; i <= jinzita2; i++ {
		for j := 0; j < jinzita2 - i; j++ {
			fmt.Print(" ")
		}
		for j := 0; j < 2 * i -1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	
	//练习四
	//打印空心金字塔
	var jinzita3 int
	fmt.Printf("输入金字塔（三）的层级：")
	fmt.Scanln(&jinzita3)
	for i := 1; i <= jinzita3; i++ {
		for j := 0; j < jinzita3 - i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2 * i - 1; j++ {
			if j == 1 || j == 2 * i - 1 || i == jinzita3 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}

	//联系五 === 99乘法表
	fmt.Println("===================================================== 99乘法表 1 =====================================================")
	var num91 int = 9
	var num92 int = 9
	for i := 1; i <= num91; i++ {
		for j := 1; j <= num92; j++ {
			if j <= i {
				fmt.Printf("%d X %d = %d   ", j, i, j * i)
			}
		}
		fmt.Println()
	}
	fmt.Println("===================================================== 99乘法表 2 =====================================================")
	for i := 1; i <= num91; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d X %d = %d \t", j, i, j * i)
		}
		fmt.Println()
	}
}
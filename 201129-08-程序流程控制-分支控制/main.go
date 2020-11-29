package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
	"math"
)

func main() {
	//顺序控制
	//分支控制
	var num1 int = 9
	if num1 == 9 {
		fmt.Println("分支1")
	} else if num1 > 9 {
		fmt.Println("分支2")
	} else {
		fmt.Println("分支3")
	}

	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scanln(&age)
	if age > 18 {
		fmt.Println("你的年龄大于18岁，请对自己负责！")
	}

	if age:= 20; age > 18 {
	 	fmt.Println("你的年龄大于18岁，请对自己负责！")
	}


	var a float64 = 3
	var b float64 = 100
	var c float64 = 6

	m := b * b - 4 * a * c

	if m > 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		x2 := (-b - math.Sqrt(m)) / 2 * a
		fmt.Println("两个解：", x1, x2)
	} else if m == 0 {
		x1 := (-b + math.Sqrt(m)) / 2 * a
		fmt.Println("一个解：", x1)
	} else {
		fmt.Println("无解...")
	}

	//1、switch 每个 case 中不需要加 break
	//2、case 后面可以有多个表达式，用 , 隔开
	var num2 int = 6
	var num3 int = 9
	switch num2 {
		case 6, 7, 8:
			fmt.Println("是6, 7, 8 其中一个")
		case num3:
			fmt.Println("是9")
		case 10:
			fmt.Println("是10")
		default:
			fmt.Println("都不匹配")
	}

	switch {
		case num2 == 6 || num2 == 7 || num2 == 8:
			fmt.Println("222是6, 7, 8 其中一个")
		case num2 == 9:
			fmt.Println("222是9")
		case num2 == 10:
			fmt.Println("222是10")
		default:
			fmt.Println("222都不匹配")
	}

	//不推荐，但是可以
	switch num3 := 10; {
		case num3 == 6 || num3 == 7 || num3 == 8:
			fmt.Println("333是6, 7, 8 其中一个")
		case num3 == 9:
			fmt.Println("333是9")
		case num3 == 10:
			fmt.Println("333是10")
		default:
			fmt.Println("333都不匹配")
	}

	//switch 穿透 fallthrough 穿透一层
	num4 := 7
	switch num4 {
		case 6,7,8:
			fmt.Println("fallthrough => 是6, 7, 8 其中一个")
			fallthrough
		case 9:
			fmt.Println("fallthrough => 是9")
		case 10:
			fmt.Println("fallthrough => 是10")
			fallthrough
		default:
			fmt.Println("fallthrough => 都不匹配")
	}

}
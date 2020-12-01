package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
	"math/rand"
	"time"
)

func main() {
	//多重循环控制 - 一般使用两层

	//TODO break
	var count int = 0
	// fmt.Println("时间戳是：", time.Now().Unix())//时间戳
	fmt.Println("时间戳是：", time.Now().UnixNano())//纳秒时间戳
	for {
		rand.Seed(time.Now().UnixNano())
		num1 := rand.Intn(100) + 1
		fmt.Println(count, "随机数是：", num1)
		if num1 == 82 {
			break
		}
		fmt.Println(count, "随机数是：", num1)
		count++
	}
	fmt.Println("用了多少次：", count)


	//break 的标签使用
	var count2 int
	ab1:
	for i := 0; i < 9; i++ {
		// ab2:
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				count2++
				fmt.Println("break:", count2)
				if k == 0 {
					break ab1
				}
			}
		}
	}

	//练习1
	var sum int
	for i := 1; i <= 100; i++ {
		sum += i
		if sum > 20 {
			fmt.Println("sum 大于20 时，当前数是：", i)
			break
		}
	}

	//练习2
	var name string
	var pass string
	for i := 0; i < 3; i++ {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Println("请输入密码：")
		fmt.Scanln(&pass)
		if name == "张无忌" && pass == "888" {
			fmt.Println("输入正确")
			break
		} else {
			fmt.Printf("输入错误，你还有%v次机会 \n", (2 - i))
		}
	}


	//TODO continue
	adb1:
	for i := 0; i < 9; i++ {
		// adb2:
		for j := 0; j < 9; j++ {
			if ( i + j ) % 2 == 0 {
				fmt.Println("i + j 是2的倍数", i + j)
			} else {
				continue adb1
			}
			fmt.Println("==========")
		}
	}

	for i := 1; i < 100; i++ {
		if i % 2 == 1 {
			fmt.Println("奇数：", i)
		} else {
			continue
		}
	}
}
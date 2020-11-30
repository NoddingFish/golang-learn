package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//循环控制
	for i := 0; i < 10; i++ {
		fmt.Println("你好，世界...", i + 1)
	}

	j := 0
	for j < 10 {
		fmt.Println("你好，世界!!!...", j + 1)
		j++
	}

	m := 0
	for {
		fmt.Println("你好，世界~~~", m + 1)
		m++
		if m > 3 {
			break
		}
	}

	//使用细节
	//TODO 遍历字符串-方法一
	//一个汉字在 utf-8 编码中占3个字节，解决：将 str 转成 []rune 切片
	var str1 string = "hello背景"
	str11 := []rune(str1) // 解决方法
	for i := 0; i < len(str11); i++ {
		fmt.Printf("%c \n", str11[i])
	}

	//TODO 遍历字符串-方法二 for-range
	//默认安装字符的方式遍历，有中文也是可以的
	var str2 string = "world你好中国"
	for index, val := range str2 {
		fmt.Printf("index=%d val=%c \n", index, val)
	}

	//练习一
	var count int = 0
	var sum int = 0
	for i := 1; i <= 100; i++ {
		if i % 9 == 0  {
			count++
			sum += i
		}
	}
	fmt.Printf("count=%d sum=%d \n", count, sum)

	//练习二
	var num1 int
	fmt.Println("请输入一个数：")
	fmt.Scanln(&num1)
	for i := 0; i <= num1; i++ {
		fmt.Printf("%d + %d = %d \n", i, num1 - i, num1)
	}

	//golang 中没有 while 和 do ... while
	
}
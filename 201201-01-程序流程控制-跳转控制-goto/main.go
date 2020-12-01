package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//TODO 多重循环控制 - goto 一般不主张使用 goto 语句
	var n int = 20
	fmt.Println("结构体1")
	if n > 10 {
		goto label1
	}
	fmt.Println("结构体2")
	label1: 
	fmt.Println("结构体3")
	fmt.Println("结构体4")

	//TODO return 一般使用再方法或者函数中
	fmt.Println(test(-2))
	fmt.Println(test(2))
}

func test(num int) string {
	if num < 0 {
		return "111111111"
	}
	return "222222222"
}
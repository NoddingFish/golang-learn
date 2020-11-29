package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//TODO 关系（比较）运算符
	fmt.Println("======================= 比较运算符 =====================")
	var n1 int = 9
	var n2 int = 8
	fmt.Println("n1>n2:", n1 > n2)
	fmt.Println("n1<=n2:", n1 <= n2)
	fmt.Println("n1==n2:", n1 == n2)
	flag := n1 > n2
	fmt.Println("flag", flag)

	//关系运算符的结果都是bool值，要么 true 要么 false
}
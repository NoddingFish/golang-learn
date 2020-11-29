package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//TODO 逻辑运算符 && || !
	fmt.Println("======================= 逻辑运算符 =====================")
	//最终结果也是 bool 值
	var num1 int = 30
	var a bool = true
	var b bool = false

	fmt.Println("a && b=", a && b)
	fmt.Println("a || b=", a || b)
	fmt.Println("!a || b=", !a || b)
	fmt.Println("!(a || b)=", !(a || b))

	if num1 > 9 {
		fmt.Println("结构体1")
	} else {
		fmt.Println("结构体2")
	}

	// && ，短路与 ，第一个条件为 false , 后面的条件不用判断
	// || ，短路或 ，第一个条件为 true , 后面的条件不用判断

}
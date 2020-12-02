package main

import (
	"fmt"
)

func main()  {
	//TODO 函数也是一种数据类型，可以赋值给一个变量，可以通过该变量对函数调用
	fmt.Println("======================================")
	a := sum
	fmt.Printf("a的类型是%T, sum的类型是%T \n", a, sum)
	fmt.Println("a的值是：", a(10, 20))

	//TODO 函数既然是一种数据类型，因为，可作为形参，并且调用
	fmt.Println("======================================")
	fmt.Println("myFun是：", myFun(sum, 100, 20))

	//TODO 自定义数据类型
	fmt.Println("======================================")
	type myInt int // myInt 和 int 虽然都是 int 类型，但是 go 认为是两个类型
	var num1 myInt = 90
	var num2 int
	num2 = int(num1) // 必须强制转换
	fmt.Println("num1=", num1)
	fmt.Println("num2=", num2)

	//案例
	fmt.Println("======================================")
	fmt.Println("myFun2是：", myFun2(sum, 100, 200))

	//案例
	fmt.Println("======================================")
	res1, res2 := test(500, 200)
	fmt.Println("test是：", res1, res2)
	
	//TODO 函数返回值可以用 _ 占位，表示忽略

	//TODO golang 支持多个参数
	//要求，编写一个 nSum 函数，可以求出1到n多个int的和
	//TODO 注意，形参列表中有可变参数，需要放在最后
	res3 := nSum(4, 7, 7, -9, 0, 10)
	fmt.Println("多个参数求和：", res3)
}

type newFun func (int, int) int

func sum(num1 int, num2 int) int {
	return num1 + num2
}

func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func myFun2(funvar newFun, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func test(num1 int, num2 int) (sum int, sub int) {
	sub = num1 - num2
	sum = num1 + num2
	return
}

func nSum(n1 int, testName... int) int {
	sum := n1
	for i := 0; i < len(testName); i++ {
		sum += testName[i]
	}
	return sum
} 
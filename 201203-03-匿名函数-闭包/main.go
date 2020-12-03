package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

var (
	//Func1 就是一个全局匿名函数 首字母大写 public
	func1 = func (n1 int, n2 int) int {
		return n1 * n2
	}
)

func AddUpper() func(int) int {
	//TODO ======= 以下是闭包
	var n int = 10
	return func (x int) int {
		n = n + x
		return n
	}
	//TODO ======= 以上是闭包
}

func main() {
	//TODO　匿名函数，没有名字的函数：如果某个函数只希望使用一次，那可以考虑使用匿名函数
	//使用方式1 - 定义匿名函数直接调用  只能用一次
	fmt.Println("========================================= 匿名函数 =========================================")
	//求两个数的和 - 匿名函数
	res1 := func (n1, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println("res1=", res1)


	//使用方式2 - 将匿名函数定义给一个变量，用该变量调用匿名函数
	//求两个数的和 - 匿名函数
	a := func (n1 int , n2 int) int {
		return n1 - n2
	}
	res2 := a(20, 55)
	res3 := a(90, 80)
	fmt.Println("res2=", res2)
	fmt.Println("res3=", res3)

	//TODO 全局匿名函数
	res4 := func1(9, 8)
	fmt.Println("res4=", res4)

	//TODO　闭包。是一个函数，和与其相关的引用环境组合的一个整体
	fmt.Println("=========================================== 闭包 ===========================================")
	//累加器
	add := AddUpper()
	fmt.Println("add1=", add(1))
	fmt.Println("add1=", add(2))
	//TODO 代码解析
	//1、AddUpper 是一个函数，返回的数据类型是 func (int) int
	//2、闭包的说明：
	//3、闭包是类，函数是操作，n是字段，函数和它使用到的n构成了闭包
	//4、
}
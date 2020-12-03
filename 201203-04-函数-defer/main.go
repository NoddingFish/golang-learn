package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func sum(num1 int, num2 int) int {
	//TODO 当执行到 defer 时，会将 defer 后面的语句压入到独立的栈中
	//TODO 当函数执行完毕后，再从defer栈，按照先入后出的方式出栈，执行
	defer fmt.Println("ok1 num1=", num1)
	defer fmt.Println("ok2 num2=", num2)

	num1++
	num2++
	res := num1 + num2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	//TODO　defer - 延时机制
	//在函数中，需要创建资源（比如：数据库连接、文件句柄、锁等），为了在函数执行完毕后，及时释放资源
	res := sum(10, 20)
	fmt.Println("res=", res)
	

	//TODO 函数的参数传递
	//1、值传递 - 值拷贝 【基本数据类型 int float bool string 数组 结构体-struct】
	//2、引用传递 - 地址拷贝【指针、slice切片、map、管道chan、interface等】
	//注意：一般引用传递效率更高
}
package main

import (
	"fmt"
)

func main() {
	//基本数据类型在内存布局
	var i int = 18
	// i 的内存地址是什么
	fmt.Println("i的地址是多少：", &i)

	//指针在内存的布局
	//1、ptr 是一个指针变量
	//2、ptr 的类型是 *int
	//3、ptr 本身的指是 &i
	var ptr *int = &i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr的地址=%v\n", &ptr)

	//获取指针类型所指向的值
	fmt.Printf("指针的值：%v\n", *ptr)

}
package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//TODO 1、大小写字母，数字和 _组成 
	var _num int
	fmt.Println(_num)

	//TODO 2、数字不能开头
	// var 2num int

	//TODO 3、Golang 中严格区分大小写
	var num int = 10
	var Num int = 20
	fmt.Println(num, Num)

	//TODO 4、标识符中不能包含空格

	//TODO 5、_ 是空标识符，用于占位

	//TODO 6、不能用系统保留关键字作为标识符 25个保留关键字
	// var break int

	fmt.Println("值类型和引用类型", _num)
}
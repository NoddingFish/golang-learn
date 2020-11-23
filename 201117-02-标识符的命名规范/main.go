package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
	// "G:/Go/Learn/201117-02-标识符的命名规范/test" // 引入包
)

func main() {
	//TODO 1.1、大小写字母，数字和 _组成 
	var _num int
	fmt.Println(_num)

	//TODO 1.2、数字不能开头
	// var 2num int

	//TODO 1.3、Golang 中严格区分大小写
	var num int = 10
	var Num int = 20
	fmt.Println(num, Num)

	//TODO 1.4、标识符中不能包含空格

	//TODO 1.5、_ 是空标识符，用于占位

	//TODO 1.6、不能用系统保留关键字作为标识符 25个保留关键字
	// var break int

	fmt.Println("值类型和引用类型", _num)


	//TODO 2.1、包名 保持 package 和目录的名字一致

	//TODO 2.2、变量名、函数名、常量名：采用驼峰法

	//TODO 2.3、如果变量名、函数名、常量名首字母大写，可以被别的包访问（公有）
	//TODO      如果变量名、函数名、常量名首字母小写，只能在本包使用（私有）
}
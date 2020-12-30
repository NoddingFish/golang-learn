package main

import (
	"fmt"
	"os"
)

func main()  {
	//TODO 文件在程序中是以流的形式操作的

	//TODO 流：数据在数据源（文件）和程序（内存）之间经历的路径
	//TODO 输入流：数据在数据源（文件）到程序（内存）的路径 -- 读文件
	//TODO 输出流：数据从程序（内存）到数据源（文件）的路径 -- 写文件

	//打开文件
	file, err := os.Open("../file/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}
	fmt.Printf("file=%v \n", file) //TODO file 是指针类型

	//关闭文件
	err2 := file.Close()
	if err2 != nil {
		fmt.Println("close file err=", err2)
	}
}

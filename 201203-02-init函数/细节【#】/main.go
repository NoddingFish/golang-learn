package main

import (
	"fmt"
	"Learn/201203-02-init函数/utils"
)

var age = test()

func test() int {
	fmt.Println("==================== 执行顺序 ==================")
	fmt.Println("全局变量定义...")
	return 90
}

func init()  {
	fmt.Println("初始化...")
}

func main()  {
	fmt.Println("main函数...")
	//TODO 如果一个文件同时包含全局变量定义，init函数和main函数，
	//TODO 则执行流程是：全局变量订单 -> init 函数 -> main 函数
	fmt.Println("======================================")
	
	fmt.Println("Age=", utils.Age, "Name=", utils.Name)
}

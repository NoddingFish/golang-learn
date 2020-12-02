package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

//init 通常是完成初始化的工作
func init()  {
	fmt.Println("这个是 init 函数")
}

func main() {
	//init 函数:每一个源文件都可以包含一个 init 函数，该函数会在 main 函数之前执行
	fmt.Println("这个是 main 函数")
}
package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func test()  {
	//使用 defer + recover 来捕获和处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
			//这里可以将错误信息发送给管理员
		}
	}()
	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func main() {
	//TODO 错误处理 - golang 不支持传统的 try ... catch
	//TODO 处理方式：defer, panic, recover
	test()
	fmt.Println("main()下面的代码...")

}
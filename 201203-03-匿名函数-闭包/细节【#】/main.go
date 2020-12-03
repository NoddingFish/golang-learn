package main

import (
	"fmt"
	"strings"
)

//闭包
func makeSuffix(suffix string) func (string) string {
	return func (name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}	
}

//传统方式
func makeSuffix2(suffix string, name string) string {
	if !strings.HasSuffix(name, suffix) {
		return name + suffix
	}
	return name
}

func main()  {
	fmt.Println("==================================== 闭包 ====================================")
	//TODO 闭包的最佳实践 syrings.HasSuffix 判断某个字符串是否包含指定的后缀

	makeSuffix := makeSuffix(".jpg")
	fmt.Println("文件名处理后：", makeSuffix("winter"))
	fmt.Println("文件名处理后：", makeSuffix("winter2.jpg"))
	fmt.Println("文件名处理后：", makeSuffix("winter2.png"))

	fmt.Println("==================================== 传统方式 ====================================")
	fmt.Println("文件名处理后：", makeSuffix2(".jpg", "winter"))
	fmt.Println("文件名处理后：", makeSuffix2(".jpg", "winter2.jpg"))
	fmt.Println("文件名处理后：", makeSuffix2(".jpg", "winter2.png"))
	
}

package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
	util "Learn/201202-03-包/utils" //支持别名
)

func main() {
	//TODO 包 golang 是以包的形式来管理项目
	var num1 float64 = 2.6
	var num2 float64 = 2.2
	res := util.Suanshu2(num1, num2, '+')
	fmt.Printf("结果是：%.3f \n", res)
	fmt.Printf("Num1是：%d \n", util.Num1)

	//注意点
	//1、一个文件打包时，该包对应一个文件夹，通常文件夹和包名一致
	//2、使用其他包的函数或变量，需要先引入该包
	//3、package 包名 第一行
	//4、路径从 $GOPATH 下的 src 开始找，不用写 src
	//5、变量名或函数名首字母 大写 可以被其他包访问
	//6、支持给包去别名
}

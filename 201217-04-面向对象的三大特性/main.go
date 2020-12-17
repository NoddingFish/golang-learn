package main

import _ "fmt"

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

func main()  {
	//TODO 三大特性：封装 继承 多态

	//封装：
	//* 就是把抽象出的字段和对字段的操作封装在一起，数据被保存在内部，程序的其他包只有通过被授权的操作（方法），才能对字段进行操作
	//实现封装的步骤：
	// 1、将结构体、字段（属性）的首字母小写
	// 2、给结构体所在包提供一个工厂模式的函数，首字母大写，类似构造函数 => 见201217-03代码
	// 3、提供一个首字母大写的Get方法，用于获取属性的值

}
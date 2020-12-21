package main

import (
	"fmt"
)

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say()  {
	fmt.Println("stu Say()")
}

type integer int

func (i integer) Say()  {
	fmt.Println("integer Say i = ", i)
}

type BInterface interface {
	// Name string//!接口中不能有任何变量
	Hello()
}

type Monster struct {

}
func (m Monster) Hello()  {
	fmt.Println("Monster Hello()")
}
func (m Monster) Say()  {
	fmt.Println("Monster Say()")
}

func main()  {
	fmt.Println("============================ 细节1 ============================")
	//TODO 1、接口本身不能创建实例，但是可以指向一个实现了该接口的自定义类型的变量（实例）
	// var a AInterface
	// a.Say()//!报错
	var stu Stu
	var a AInterface = stu
	a.Say()

	//TODO 2、接口中的所有方法都没有方法体，即都是没有实现的方法

	//TODO 3、自定义类型将某个接口的所有方法都实现，才说实现了该接口

	//TODO 4、自定义类型只有实现了某个接口，才能将该自定义类型的实例赋值给接口类型

	//TODO 5、自定义类型，不一定是 struct
	var i integer = 10
	var b AInterface = i
	b.Say()

	//TODO 6、一个自定义类型可以实现多个接口
	var mon Monster
	var a6 AInterface = mon
	var b6 BInterface = mon
	a6.Say()
	b6.Hello()

	//TODO 7、接口中不能有任何变量
}
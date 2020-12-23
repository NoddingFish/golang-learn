package main

import (
	"fmt"
)

//猴子结构体
type Monkey struct {
	Name string
}

func (this Monkey) climbing()  {
	fmt.Println(this.Name, " 生来会爬树...")
}

//小猴子继承了猴子的结构体
type LittleMonkey struct {
	Monkey
}

//声明接口
type BirdAble interface{
	Flying()
}
type FishAble interface{
	Swimming()
}

func (liMon LittleMonkey) Flying()  {
	fmt.Println(liMon.Name, " 学会了飞翔...")
}
func (liMon *LittleMonkey) Swimming()  {
	fmt.Println(liMon.Name, " 学会了游泳...")
}

func main()  {
	//TODO 接口是对继承的一种补充
	monkey := LittleMonkey{
		Monkey{
			Name: "悟空",
		},
	}
	monkey.climbing()
	monkey.Flying()
	monkey.Swimming()

	//TODO 2、继承：解决代码的复用性和可维护性
	//TODO    接口：设计、设计好各种规范（方法）

	//TODO 3、接口比继承更加灵活

	//TODO 4、接口在一定程度上实现代码解耦
}
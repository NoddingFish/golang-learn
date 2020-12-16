package main

import (
	"fmt"
)
type Circle struct {
	radius float64
}

type interger int

func (i interger) print()  {
	fmt.Println("print i=", i)
}

func (i *interger) change()  {
	*i += 1
}

func (c *Circle) area() float64 {
	c.radius = 10
	// return 3.14 * (*c).radius * (*c).radius//标准的访问是
	return 3.14 * c.radius * c.radius//!优化后的写法
}

type Student struct {
	Name string
	Age int
}

func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

func main()  {
	//TODO 1、结构体是值类型
	//TODO 2、修改结构体的值，可以通过结构体指针的方式
	var cir Circle
	cir.radius = 4
	// res := (&cir).area()//标准写法应该是这样
	res := cir.area()//!优化后的写法
	fmt.Println("面积是=", res)
	fmt.Println("radius=", cir.radius)

	//TODO 3、方法作用再指定的数据类型上，不仅仅是struct，int等都可以
	var i interger = 10
	i.print()
	i.change()
	fmt.Println("change i=", i)

	//TODO 4、方法名首字母大小写，大写可以在其他包访问

	//TODO 5、如果一个类型实现了 String() 这个方法，那么fmt.Println 默认会调用这个变量的 String() 进行输出
	fmt.Println("=========================== String() ===========================")
	stu := Student{
		Name: "tom",
		Age: 10,
	}
	fmt.Println(&stu)
}
package main

import (
	"fmt"
)

type Person struct {
	Name string
}

//TODO func (a A) test() {} 表示A结构体有一个方法，方法名为 test()
//TODO (a A)
func (p Person) test() {
	p.Name = "jack"
	fmt.Println("test() = ", p.Name)
}

func main()  {
	//TODO 方法是作用再指定的数据类型上的，和指定的数据类型绑定
	//TODO 因此自定义类型，都可以有方法，不仅仅是 struct

	var person Person
	person.Name = "tom"
	person.test()//调用方法
	fmt.Println("person = ", person.Name)
}
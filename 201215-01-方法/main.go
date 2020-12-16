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

func (p Person) speak() {
	fmt.Println(p.Name, "是一个好人")
}

func (p Person) jiSu() int {
	count := 0
	for i := 1; i <= 1000; i++ {
		count += i
	}
	return count
}

func (p Person) jiSu2(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		count += i
	}
	return count
}

func (p Person) jiSu3(n1, n2 int) int {
	return n1 + n2
}

func main()  {
	//TODO 方法是作用再指定的数据类型上的，和指定的数据类型绑定
	//TODO 因此自定义类型，都可以有方法，不仅仅是 struct

	var person Person
	person.Name = "tom"
	person.test()//调用方法
	fmt.Println("person = ", person.Name)

	person.speak()

	count := person.jiSu()
	fmt.Println("count = ", count)

	count2 := person.jiSu2(1000)
	fmt.Println("count2 = ", count2)

	count3 := person.jiSu3(10, 20)
	fmt.Println("count3 = ", count3)
}
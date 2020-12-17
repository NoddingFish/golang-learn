package main

import (
	"fmt"
)

type Person struct {
	Name string
}

func test01(p Person)  {
	fmt.Println("test01 = ", p.Name)
}

func test02(p *Person)  {
	fmt.Println("test02 = ", p.Name)
}

func (person Person) test03()  {
	person.Name = "jack" // person 还是值拷贝
	fmt.Println("test03 = ", person.Name)
}

func (person *Person) test04()  {
	person.Name = "mary" // person 还是值拷贝
	fmt.Println("test04 = ", person.Name)
}

func main()  {
	//TODO 调用方式不一样 函数：直接调用 方法：变量.方法()

	//TODO 
	p := Person{Name: "tom"}
	test01(p)//不用使用 test01(&p)
	test02(&p)//不用使用 test01(p)

	p.test03()
	fmt.Println("p.Name = ", p.Name)
	(&p).test03() //! 虽然是 &p，但是里面还是值拷贝
	fmt.Println("p.Name = ", p.Name)


	p.test04()//都可以
	fmt.Println("p.Name = ", p.Name)
	(&p).test04()//都可以
	fmt.Println("p.Name = ", p.Name)

	//TODO 总结：
	//TODO 不过调用形式如何，真正决定是值拷贝还是地址拷贝，看这个方法是和哪个类型绑定
	//TODO 如果是和值类型，比如 (p Person)，则是值拷贝，如果是和指针类型，比如（p *Person）则是地址拷贝
}
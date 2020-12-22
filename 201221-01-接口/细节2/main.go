package main

import (
	"fmt"
)

type BInterface interface {
	test01()
}

type CInterface interface {
	test02()
}

type AInterface interface {
	BInterface
	CInterface
	test03()
}

type Stu struct {

}

func (stu Stu) test01()  {
	fmt.Println("test01()")
}

func (stu Stu) test02()  {
	fmt.Println("test02()")
}

func (stu Stu) test03()  {
	fmt.Println("test03()")
}

// ================= 空接口
type T interface {

}

func main()  {
	fmt.Println("============================ 细节2 ============================")
	//TODO 8、一个接口（A接口）可以继承多个别的接口（B和C），如果要实现A接口，那必须将B、C接口的方法都实现
	var stu Stu
	var a AInterface = stu
	var b BInterface = stu
	a.test01()
	a.test02()
	b.test01()
	// b.test02()//!报错 b接口没有 test02方法

	//TODO 9、interface 默认是一个指针（引用类型）

	//TODO 10、空接口 interface{} 没有任何方法，所以所有类型都实现了空接口
	//TODO 可以给任何变量赋给空接口
	var t T = stu
	fmt.Println("空接口：", t)
	var t2 interface{} = stu
	var num1 float64 = 8.8
	t2 = num1
	fmt.Println("空接口t2：", t2)
}
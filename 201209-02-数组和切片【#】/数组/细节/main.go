package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func test01(arr [3]int)  {
	arr[0] = 88
}

func test02(arr *[3]int)  {
	// arr[0] = 88
	(*arr)[0] = 88 //*arr 取arr的地址的指针
}

func main() {
	//* 数组的细节
	//TODO! 1、数组是多个相同类型的数据的组合，一个数组一旦声明/定义了，长度是固定的，不能动态变化的
	// var arr01 [3]int
	arr01 := [...]int{2,3,4,5}
	fmt.Println("arr01的长度=", len(arr01))
	fmt.Printf("arr01的类型=%T \n", arr01)

	//TODO! 2、var arr []int 这时 arr 是个切片 slice

	//TODO 3、数组中的原色可以是任何数据类型，包括值类型和引用类型，但是不能混用
	var arr02 [3]*int
	num1 := 7
	arr02[0] = &num1
	arr02[1] = &num1
	arr02[2] = &num1
	fmt.Println("arr01是=", arr02)

	//TODO 4、数组创建后，有默认值 int=>0 string=>"" bool=>false

	//TODO 5、数组的使用步骤：1、声明数组并开辟空间 2、赋值 3、使用数组

	//TODO 6、数组的下标识从 0 开始

	//TODO! 7、数组属于值类型，默认情况下是值传递，会进行值拷贝，数组间不会相互影响
	arr07 := [3]int{11, 22, 33}
	test01(arr07)
	fmt.Println("arr07=", arr07)

	//TODO 8、如果在其他函数中，去修改原来的数组，可以使用引用传递【指针类型】
	arr08 := [3]int{11, 22, 33}
	test02(&arr08)
	fmt.Println("arr08=", arr08)

	//TODO 9、长度是数据类型的一部分， [3]int 和 [4]int 不是同一类型
}
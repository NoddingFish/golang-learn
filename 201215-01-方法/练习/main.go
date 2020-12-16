package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

type Meth struct {
	//可以空
}

func (mu Meth) Print()  {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (mu Meth) JudeNum(num int)  {
	if num % 2 == 0 {
		fmt.Println(num, "是偶数")
	} else {
		fmt.Println(num, "是奇数")
	}
}

func main()  {
	fmt.Println("=========================== 练习1 ===========================")
	var cir Circle
	cir.radius = 4
	res := cir.area()
	fmt.Println("面积是=", res)

	fmt.Println("=========================== 练习2 ===========================")
	var mu Meth
	mu.Print()

	fmt.Println("=========================== 练习2 ===========================")
	mu.JudeNum(11)
}
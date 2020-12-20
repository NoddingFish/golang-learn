package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	a A //TODO 有名结构体
}

type Goods struct {
	Name string
	Price int
}

type Brand struct {
	Name string
	Address string
}

type TV struct {
	Goods
	*Brand
}

func main()  {
	//TODO 细节5、组合：如果B中是有名结构体是，那访问，就必须得带上有名结构体的名字
	var b B
	// b.Name = "tom"//!报错
	b.a.Name = "mary~"
	fmt.Println("b=", b)

	//TODO 细节6、
	tv := TV{
		Goods{"电视0001",5000},
		&Brand{Address:"山东", Name:"海尔"},
	}
	fmt.Println("tv=", tv)
	fmt.Println("tv=", *tv.Brand)
}
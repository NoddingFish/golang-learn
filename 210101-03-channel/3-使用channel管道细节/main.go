package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age int
}

func main()  {
	allChan := make(chan interface{}, 3)

	allChan<- 11
	allChan<- "tom jack"
	cat := Cat{"小花猫", 10}
	allChan<- cat

	<-allChan
	<-allChan

	newCat := <-allChan //这是什么？

	fmt.Printf("newCat=%T, newCat=%v \n", newCat, newCat)
	// fmt.Printf("newCat.Name=%v \n", newCat.Name)//!报错 编译出错

	//解决办法：类型断言
	a := newCat.(Cat)
	fmt.Printf("newCat.Name=%v \n", a.Name)
}
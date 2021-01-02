package main

import (
	"fmt"
)

func main()  {
	
	intChan := make(chan int, 3)
	intChan<- 100
	intChan<- 200

	close(intChan)//TODO 内建函数，关闭管道
	//这时不能再写入数据到 intChan,但是读取是可以的
	// intChan<- 300
	fmt.Println("ok~")

	n1 := <-intChan
	fmt.Println("n1=", n1)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2<- i * 2
	}

	//TODO 需要关闭管道，若没有关闭会报错 deadlock，但是数据都可以取出来
	close(intChan2)
	
	//! 管道只有一个返回值，没有下标
	for v := range intChan2 {
		fmt.Println("v=", v)
	}
}
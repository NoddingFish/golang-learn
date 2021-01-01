package main

import (
	"fmt"
)


//TODO 1、channel 本质就是一个数据结构-队列
//TODO 2、先进先出【FIFO】
//TODO 3、线程安全，多 goroutine 访问，不需要加锁
//TODO 4、channel 有类型的，一个 string 类型的 channel 只能放 string 类型的数据

// channel 是引用类型
// channel 必须初始化才能写入数据 make()



func main()  {
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Println("intChan 是：", intChan)

	intChan<- 11 //写入符号 <-
	// intChan<- 11.9 //! 必须是 int 类型
	num := 22
	intChan <- num
	intChan <- 33
	// intChan <- 44//!不能超过 intChan 的容量

	<-intChan //! 可以直接取，不使用，取出后，可以放入新的数据
	intChan <- 55

	fmt.Printf("intChan的长度是：%v 容量是：%v \n", len(intChan), cap(intChan))

	//但是可以取
	num1 := <-intChan
	fmt.Println("num1=", num1)
	fmt.Printf("intChan的长度是：%v 容量是：%v \n", len(intChan), cap(intChan))

	//TODO 在没有使用协程的情况下，如果我们的管道数据全部取出，再取就会报错 deadlock
}
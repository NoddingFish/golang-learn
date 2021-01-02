package main

import (
	"fmt"
	"time"
)

func WriteData(intChan chan int)  {
	for i := 1; i <= 50; i++ {
		intChan<- i
		fmt.Println("WriteData 写入数据=", i)
		// time.Sleep(time.Second)
	}
	close(intChan)//关闭
}

func ReadData(intChan chan int, exitChan chan bool)  {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("ReadData 读到数据=", v)
		time.Sleep(time.Second)//TODO 写的很快，但是读的很慢，频率不一致，无所谓
	}
	//读取完成后
	exitChan<- true
	close(exitChan)
}

func main()  {
	intChan := make(chan int, 10)
	exitChan := make(chan bool, 1)

	go WriteData(intChan)
	go ReadData(intChan, exitChan)

	//传统
	// time.Sleep(time.Second * 2)

	for {
		_, ok := <-exitChan
		if !ok {
			break
		}
	}
}
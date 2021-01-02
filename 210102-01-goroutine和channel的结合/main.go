package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int)  {
	for i := 1; i <= 200000; i++ {
		intChan<- i
	}
	close(intChan)//关闭 intChan
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool)  {
	var flag bool
	for {
		// time.Sleep(time.Millisecond * 100)
		num, ok := <-intChan
		if !ok {
			break
		}
		flag = true
		for i := 2; i< num ; i++ {
			if num % i == 0 {
				//说明不是素数
				flag = false
				break
			}
		}
		if flag {
			//是素数
			primeChan<- num
		}
	}
	fmt.Println("有一个 primeChan 协程因为取不到数据，退出..")
	//这里不能关闭 primeChan
	exitChan<- true
}

func main()  {
	start := time.Now().Unix()

	intChan := make(chan int , 1000)

	primeChan := make(chan int, 2000)

	exitChan := make(chan bool, 4)

	//开启一个协程
	go putNum(intChan)

	//开启4个协程，从 intChan 取数据，判断是否是素数
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//处理：
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		end := time.Now().Unix()

		fmt.Println("使用协程耗时：", end - start)
		close(primeChan)
	}()

	for {
		res, ok := <- primeChan
		if !ok {
			break
		}
		fmt.Println("素数=", res)
	}

	fmt.Println("主线程退出!")
}
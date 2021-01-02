package main

import (
	"fmt"
)

func putNum(numChan chan int) {
	for i := 1; i <= 20; i++ {
		numChan<- i
	}
	close(numChan)
}

func sumNum(numChan chan int, sumChan chan map[int]int,exitChan chan bool) {
	for {
		num, ok := <-numChan
		if !ok {
			break
		}
		sumChan<- getSum(num)
	}
	fmt.Println("有一个协程已完成....")
	exitChan<- true
}

func getSum(n int) map[int]int {
	res := make(map[int]int, 1)
	num := 0
	for i := 1; i <= n; i++ {
		num += i
	}
	res[n] = num
	return res
}

func main()  {
	numChan := make(chan int, 1000)
	sumChan := make(chan map[int]int, 1000)
	exitChan := make(chan bool, 8)
	go putNum(numChan)

	for i := 1; i <= 8; i++ {
		go sumNum(numChan, sumChan, exitChan)
	}

	go func ()  {
		for i := 0; i < 8; i++ {
			<-exitChan
		}
		close(sumChan)
	}()

	for {
		sum , ok := <-sumChan
		if !ok {
			break
		}
		fmt.Printf("sum=%v \n", sum)
	}

	fmt.Println("主线程退出...")
}
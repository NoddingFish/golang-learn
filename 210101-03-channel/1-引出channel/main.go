package main

import (
	"fmt"
	"time"
)

var (
	myMap = make(map[int]int, 10)
)

func test(n int)  {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	myMap[n] = res
}

func main()  {
	for i := 1; i <= 200; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 5)

	for i, v := range myMap {
		fmt.Printf("myMap[%d]=%d \n", i, v)
	}

	fmt.Println("结束...")

	//TODO 运行程序，如何知道是否存在资源竞争问题， ==> 增加 -race 参数
	//TODO 列如： go build -race main.go ==> Found 3 data race(s)
}
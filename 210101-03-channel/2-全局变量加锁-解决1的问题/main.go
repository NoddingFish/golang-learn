package main

import (
	"fmt"
	"time"
	"sync"
)

var (
	myMap = make(map[int]int, 10)

	//TODO 声明一个全局的互斥锁
	//Mutex 互斥
	lock sync.Mutex
)

func test(n int)  {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	lock.Lock()//!加锁
	myMap[n] = res
	lock.Unlock()//!解锁
}

func main()  {
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	time.Sleep(time.Second * 5)//!必须指定时间，不好，再改进

	lock.Lock()//!加锁
	for i, v := range myMap {
		fmt.Printf("myMap[%d]=%d \n", i, v)
	}
	lock.Unlock()//!解锁

	fmt.Println("结束...")

	//TODO 运行程序，如何知道是否存在资源竞争问题， ==> 增加 -race 参数
	//TODO 列如： go build -race main.go ==> Found 3 data race(s)
}
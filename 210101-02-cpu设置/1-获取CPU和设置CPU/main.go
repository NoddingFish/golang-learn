package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置使用多少个CPU
	runtime.GOMAXPROCS(cpuNum - 1)//TODO go1.8之后可以不用设置
	fmt.Println("ok")
}

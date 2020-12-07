package main

import (
	"fmt"
	"time"
	"strconv"
)

func test01()  {
	str := ""
	for i := 0; i < 1000; i++ {
		str += "hello" + strconv.Itoa(i) //整数转字符串
	}
}

func main()  {
	// now := time.Now()
	// start := now.Unix()
	// test01()
	// end := now.Unix()
	//以上是错误写法
	start := time.Now().UnixNano()
	test01()
	end := time.Now().UnixNano()
	fmt.Println("运行程序 test01 耗时：", end - start, "纳秒")

}
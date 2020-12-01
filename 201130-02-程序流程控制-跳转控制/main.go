package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
	"math/rand"
	"time"
)

func main() {
	//多重循环控制 - 一般使用两层
	var count int = 0
	fmt.Println("时间戳是：", time.Now().Unix())//时间戳
	for {
		rand.Seed(time.Now().Unix())
		num1 := rand.Intn(100) + 1
		fmt.Println(count, "随机数是：", num1)
		if num1 == 82 {
			break
		}
		fmt.Println(count, "随机数是：", num1)
		count++
	}
	fmt.Println("用了多少次：", count)
}
package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
	"time"
)

func main() {
	//TODO　日期和时间相关的函数 - 需要导入 time 包

	//TODO 1、获取当前时间
	now := time.Now()
	fmt.Printf("获取当前时间：%v 类型是：%T \n", now, now)
	
	//TODO 2、获取其他的日期信息
	fmt.Println("获取当前时间的年份：", now.Year())
	fmt.Println("获取当前时间的月份：", now.Month())
	fmt.Println("获取当前时间的月份：", int(now.Month()))
	fmt.Println("获取当前时间的日期：", now.Day())
	fmt.Println("获取当前时间的时间：", now.Hour())
	fmt.Println("获取当前时间的分钟：", now.Minute())
	fmt.Println("获取当前时间的秒钟：", now.Second())

	//TODO 2、格式化日期时间 - 可自由组合 - 数字是固定的
	fmt.Println("格式化日期时间：", now.Format("2006/01/02 15:04:05"))
	fmt.Println("格式化日期：", now.Format("2006-01-02"))
	fmt.Println("格式化时间：", now.Format("15:04:05"))

	//TODO 3、时间常量
	// const (
	// 	Nanosecond  Duration = 1
	// 	Microsecond          = 1000 * Nanosecond
	// 	Millisecond          = 1000 * Microsecond
	// 	Second               = 1000 * Millisecond
	// 	Minute               = 60 * Second
	// 	Hour                 = 60 * Minute
	// )
	i := 0
	for {
		i++
		fmt.Println("i=", i)
		// time.Sleep(time.Second)//休眠 1秒
		// time.Sleep(time.Millisecond * 100)//休眠 0.1秒
		time.Sleep(time.Nanosecond)//休眠 0.1纳秒
		if (i == 10) {
			break
		}
	}

	//TODO 4、Unix时间戳 Unixnano时间戳
	fmt.Println("===================================== Unix时间戳 unixnano时间戳 ===============================")
	fmt.Println("时间戳", now.Unix())
	fmt.Println("纳秒时间戳", now.UnixNano())
}
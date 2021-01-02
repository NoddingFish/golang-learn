//TODO 细节说明：见 markdown 文档

package main

import (
	"fmt"
	"time"
)

func sayHello()  {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("hello world")
	}
}

func test()  {
	//可以捕获到异常
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误：", err)
		}
	}()
	var myMap map[int]string
	myMap[0] = "golang" //会报错 panic ，因为需要先 make
}

func main() {
	go sayHello()
	go test() // 会报错 panic，导致程序崩溃

	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)
		fmt.Println("main() ok=", i)
	}
}



// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	intChan := make(chan int, 10)
// 	for i := 0; i < 10; i++ {
// 		intChan <- i
// 	}

// 	strChan := make(chan string, 10)
// 	for i := 0; i < 10; i++ {
// 		strChan <- "hello " + fmt.Sprintf("%d", i)
// 	}

// 	// label://退出程序-方式2
// 	for {
// 		select {
// 			case v := <- intChan :
// 				fmt.Printf("从 intChan 读取的数据： %d \n", v)
// 				time.Sleep(time.Millisecond * 100)
// 			case v := <- strChan :
// 				fmt.Printf("从 strChan 读取的数据： %s \n", v)
// 				time.Sleep(time.Millisecond * 100)
// 			default:
// 				fmt.Printf("都取不到了，不玩了，程序员可以加入自己的逻辑")
// 				time.Sleep(time.Millisecond * 100)
// 				return //退出程序-方式1 【推荐】
// 				// break label//退出程序-方式2
// 		}
// 	}
// }
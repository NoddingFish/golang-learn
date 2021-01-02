## goroutine 和 channel 的细节：

1. 管道可以声明为只读或者只写

   1. 只写：

      ```go
      var chan2 chan<- int
      chan2 = make(chan int, 3)
      chan2<- 20
      num := <- chan2//这里会报错
      fmt.Println("chan2=", chen2)
      ```

   2. 只读：

      ```go
      var chan3 <-chan int
      num2 := <- chan3
      chan3<- 30//这里会报错
      fmt.Println("num2=", num2)
      ```

      > 注意：
      >
      > 只读 `<-chan` 只写 `chan<-` 可以使用在防止误操作的地方
      >
      > `chan int` 类型 和 `<-chan int` 和 `chan<- int` 本质上还是 `chan int` 类型，可以通用

      

2. 使用 `select` 可以解决从管道取数据的阻塞问题

   在实际开发中，可能不好确定什么时候关闭管道，可以使用 `select` 方式解决

   ```go
   package main
   
   import (
   	"fmt"
   	"time"
   )
   
   func main() {
   	intChan := make(chan int, 10)
   	for i := 0; i < 10; i++ {
   		intChan <- i
   	}
   
   	strChan := make(chan string, 10)
   	for i := 0; i < 10; i++ {
   		strChan <- "hello " + fmt.Sprintf("%d", i)
   	}
   
   	// label://退出程序-方式2
   	for {
   		select {
   			case v := <- intChan :
   				fmt.Printf("从 intChan 读取的数据： %d \n", v)
   				time.Sleep(time.Millisecond * 100)
   			case v := <- strChan :
   				fmt.Printf("从 strChan 读取的数据： %s \n", v)
   				time.Sleep(time.Millisecond * 100)
   			default:
   				fmt.Printf("都取不到了，不玩了，程序员可以加入自己的逻辑")
   				time.Sleep(time.Millisecond * 100)
   				return //退出程序-方式1 【推荐】
   				// break label//退出程序-方式2
   		}
   	}
   }
   ```

   

3. `goroutine` 中使用 `recover` ，解决协程中出现的 `panic` ，导致程序崩溃的问题

   下面代码会报错 `panic` ，程序崩溃 `panic: assignment to entry in nil map`

   ```go
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
   	//定义map - 错误写法
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
   ```

   可以使用 `defer` 和 `recover` 解决：

   ```go
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
   	//重点
   	defer func() {
   		if err := recover(); err != nil {
               //可以捕获到异常
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
   ```

   


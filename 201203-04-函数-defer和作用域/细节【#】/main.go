package main

import (
	"fmt"
)

var age int = 89
// name := "tom" //TODO 报错 =》 var name string 和 name = "tom" ，赋值语句只能用于函数内

func MysqlDb()  {
	// connect = openDatabase()//连接数据库
	// defer connect.close()
}

func main()  {
	fmt.Println("==================================== defer ====================================")
	//TODO 上面是示意代码，连接数据库可以立马用 defer 销毁
	//程序会在程序执行完成后，去主动的销毁，不需要程序员考虑什么时候销毁

	//TODO  变量作用域
}

package main

import (
	"fmt"
	"flag"
)

func main()  {
	//mysql -u root -h 127.0.0.5 -pwd 123 -p 3306
	var user string
	var pwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名")
	flag.StringVar(&pwd, "pwd", "", "密码")
	flag.StringVar(&host, "h", "localhost", "主机")
	flag.IntVar(&port, "p", 3306, "端口")

	//TODO 非常重要，需要使用 flag.Parse()
	flag.Parse()

	fmt.Printf("user=%v pwd=%v host=%v port=%v", 
	user, pwd, host, port)

	//* 执行：
	//go build -o main.exe .\test.go
	//.\main.exe -u root -pwd 222 -p 8080
}
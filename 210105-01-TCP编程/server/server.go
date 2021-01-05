package main

import (
	"fmt"
	"net"
)

func main()  {
	fmt.Println("服务器开始监听...")

	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("listen err ", err)
	}
	defer listen.Close()//延时关闭 listen

	//循环等待客户端链接
	for {
		fmt.Println("等待客户端来链接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Accrpt() err ", err)
		}
		fmt.Println("conn=", conn)
	}
	// fmt.Println("listen=", listen)
}
package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	//这里循环的接受客户端发送过来的数据
	defer conn.Close()// 关闭

	for {
		//创建一个新的切片
		buf := make([]byte, 1024)

		fmt.Printf("服务器在等待客户端 %s 发送信息\n", conn.RemoteAddr().String())
		//1、等待客户端通过 conn 发送消息
		//2、如果客户端没有 write-发送，那么协程会阻塞在这里
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("服务器的Read err ", err)
			return 
		}
		//3、显示客户端发送的内容到服务器的终端
		fmt.Print(string(buf[:n]))//TODO 重点
	}
}

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
		fmt.Printf("Accrpt() conn=%v 客户端ip:%v \n", conn, conn.RemoteAddr())

		go process(conn)
	}
	// fmt.Println("listen=", listen)
}
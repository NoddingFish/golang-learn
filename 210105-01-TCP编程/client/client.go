package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func main()  {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("client dial err ", err)
		return
	}
	// fmt.Println("conn ", conn)

	reader := bufio.NewReader(os.Stdin)//os.Stdin 代表标准输入【终端】

	for {
		//从终端读取一行，发送给服务器
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readerstring err ", err)
		}
		//处理 line
		line = strings.Trim(line, "\r\n")
		if line == "exit" {
			fmt.Println("客户端退出")
			break
		}

		n, err := conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("conn.Write err ", err)
		}
		fmt.Printf("客户端发送了 %d 字节的数据", n)
	}
}
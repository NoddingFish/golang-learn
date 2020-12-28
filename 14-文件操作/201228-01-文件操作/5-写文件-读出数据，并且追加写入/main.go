package main

import (
	"fmt"
	"bufio"
	"os"
	"io"
)

func main()  {
	//创建文件，写入 hello Gardon
	filePath := "../file/abc.txt"
	file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("open file err=%v \n", err)
		return
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//显示到终端
		fmt.Print(str)
	}

	str := "你好，东大门~  \r\n"
	//写入是，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}

	//TODO 因为 writer 是带缓存的，因此在调用 WriterString方法时，其实内容是先写入
	//TODO 到缓存的，所以需要调用 Flush 方法，将缓存数据写入到文件中
	writer.Flush()

	fmt.Println("写入文件成功...")
}

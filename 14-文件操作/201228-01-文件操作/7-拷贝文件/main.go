package main

import (
	"fmt"
	"os"
	"io"
	"bufio"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open file err=", err)
	}
	reader := bufio.NewReader(srcFile)
	
	defer srcFile.Close()

	// dstFileName 可能不存在
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file err=", err)
	}

	writer := bufio.NewWriter(dstFile)

	defer dstFile.Close()

	return io.Copy(writer, reader)
}

func main()  {
	srcFileName := "../file/avatar.jpg"
	dstFileName := "../file/avatar2.jpg"
	_, err := CopyFile(dstFileName, srcFileName)
	if err == nil {
		fmt.Println("复制完成...")
	} else {
		fmt.Println("复制错误：", err)
	}
}

package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"io/ioutil"
)

func main()  {
	fmt.Println("=============== 方式一-带缓冲 ===================")

	//打开文件
	file, err := os.Open("../file/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时的关闭file，否则会有内存泄漏
	defer file.Close()

	//创建一个 *Reader ，是带缓冲的
	reader := bufio.NewReader(file)
	//循环的读取文件的内容
	for {
		str, err := reader.ReadString('\n')//! 读到一个换行就结束
		//输出内容
		fmt.Print(str)
		if err == io.EOF {//! io.EOF 表示文件的末尾
			break
		}
	}

	fmt.Println("=============== 文件读取结束 ===================")
	
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println("=============== 方式二-一次性读取 ===================")
	//TODO 适用于不大的文件
	//使用 ioutil.ReadFile 一次性读取

	fileUrl := "../file/test.txt"
	content, err2 := ioutil.ReadFile(fileUrl)//!不需要打开关闭，已被封装在一起
	if err2 != nil {
		fmt.Printf("read file err=%v", err)
	}
	// fmt.Printf("%v", content)
	fmt.Printf("%v", string(content))
}

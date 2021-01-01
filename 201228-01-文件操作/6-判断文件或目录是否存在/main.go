package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error)  {
	_, err := os.Stat(path)
	//TODO 如果返回的错误是 nil，说明文件或文件夹存在
	if err == nil {
		return true, nil
	}
	//TODO 如果返回的错误类型使用 os.IsNotExist 判断为 true，说明文件或文件夹不存在
	if os.IsNotExist(err) {
		return false, nil
	}
	//TODO 如果返回的错误是其他类型，则不确定是否存在
	return false, err
}

func main()  {
	//创建文件，写入 hello Gardon
	isEx, err := PathExists("../file/abc.txt")
	if isEx {
		fmt.Println("文件存在...")
	} else {
		fmt.Println("文件不存在...", err)
	}
}

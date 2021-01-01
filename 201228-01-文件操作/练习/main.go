package main

import (
	"fmt"
	"io/ioutil"
)

func main()  {
	//1、将 file1.txt 文件内容导入到 file2.txt
	filePath1 := "../file/file1.txt"
	filePath2 := "../file/file2.txt"

	data, err := ioutil.ReadFile(filePath1)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
		return
	}

	err = ioutil.WriteFile(filePath2, data, 0666)
	if err != nil {
		fmt.Printf("write file err=%v\n", err)
	}
}
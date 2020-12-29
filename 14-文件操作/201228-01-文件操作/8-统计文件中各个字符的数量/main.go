package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

type CharCount struct {
	chCount int
	numCount int
	spaceCount int
	otherCount int
}

func main()  {
	//打开一个文件
	file, err := os.Open("../file/8.txt")
	if err != nil {
		fmt.Println("open file err=", err)
		return
	}

	defer file.Close()

	var count CharCount

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		for _, v := range []rune(str) {
			
			switch {
				case v >= 'a' && v <= 'z':
					fallthrough // 穿透
				case v >= 'A' && v <= 'Z':
					count.chCount++
				case v == ' ' || v == '\t':
					count.spaceCount++
				case v >= '0' && v <= '9':
					count.numCount++
				default:
					count.otherCount++
			}
		}
	}

	fmt.Println("数量：", count)

}
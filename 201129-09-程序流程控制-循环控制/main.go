package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//循环控制
	for i := 0; i < 10; i++ {
		fmt.Println("你好，世界...", i + 1)
	}

	j := 0
	for j < 10 {
		fmt.Println("你好，世界!!!...", j + 1)
		j++
	}

	m := 0
	for {
		fmt.Println("你好，世界~~~", m + 1)
		m++
		if m > 3 {
			break
		}
	}
}
package main

// import "fmt"
// import "unsafe"

import (
	"fmt"
	"unsafe"
)

func main() {

	var i int = 100

	fmt.Println(i)

	var i1 int8 = -128 //int8 -128~127

	fmt.Println(i1)

	var ui1 uint8 = 255

	fmt.Println(ui1)

	//int 的其他类型 （int uint rune byte）
	var a int = 8900
	fmt.Println(a)

	var b uint = 1
	fmt.Println(b)

	var c rune = 120
	fmt.Println(c)

	var d byte = 255
	fmt.Println(d)

	//查看变量的类型
	var i2 = 100
	fmt.Printf("i2 的类型：%T \n", i2)

	//查看变量占用的字节大小和数据类型（使用较多）
	var n2 int64 = 10
	fmt.Printf("n2 的类型 %T n2 占用的字节数是 %d", n2, unsafe.Sizeof(n2))
}

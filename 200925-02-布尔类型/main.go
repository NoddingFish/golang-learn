package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b1 bool = true

	fmt.Println("b1=", b1)
	fmt.Println("b1 的空间占用:", unsafe.Sizeof(b1))

}
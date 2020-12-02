package utils

import "fmt"

var Age int
var Name string

func init()  {
	fmt.Println("utils 中的 init 函数...")
	Age = 99
	Name = "Tom"
}
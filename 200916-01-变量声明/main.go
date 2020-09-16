package main

import "fmt"

//定义全局变量1
// var x1 = 100
// var x2 = 200
// var x_name = "jon"

//定义全局变量2
var (
	x1     = 100
	x2     = 200
	x_name = "jon"
)

func main() {
	var n1, n2, n3 int

	fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	var name, age, address = "tom", 18, "浙江省"

	fmt.Println("name=", name, "age=", age, "address=", address)

	//一次性声明多个变量方式3
	name1, age1, address1 := "tom", 18, "浙江省"

	fmt.Println("name1=", name1, "age1=", age1, "address1=", address1)

	fmt.Println("x1=", x1, "x2=", x2, "x_name=", x_name)

	var i = 10

	// i = "tom"

	fmt.Println("i=", i)

	
	var str string

	fmt.Println("str=", str)

}

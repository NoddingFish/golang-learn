package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
	"strconv"
)

func main() {
	var str1 string  ="true"
	var b bool

	b, _ = strconv.ParseBool(str1)
	fmt.Printf("b type %T  b=%v\n", b, b)

	var str2 string = "13133"
	var n1 int64
	n1, _ = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("b type %T  b=%v\n", n1, n1)

	var str3 string = "132.1112"
	var f1 float64
	f1, _ = strconv.ParseFloat(str3, 64)
	fmt.Printf("b type %T  b=%v\n", f1, f1)

	//注意事项：
	var str4 string = "2hello"
	var n2 int64
	n2, _ = strconv.ParseInt(str4, 10, 64)
	fmt.Printf("b type %T  b=%v\n", n2, n2)

	var str5 string = "2hello"
	var b2 bool
	b2, _ = strconv.ParseBool(str5)
	fmt.Printf("b type %T  b=%v\n", b2, b2)
}
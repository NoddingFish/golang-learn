package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//TODO Go 在不同类型的变量之间赋值时需要显示转换，不能自动转换

	var a1 int32 = 100
	var b1 float32 = float32(a1)
	var b2 int8 = int8(b1)

	fmt.Printf("a1=%v b1=%v b2=%v\n", a1, b1, b2)
	fmt.Printf("a1=%T b1=%T b2=%T\n", a1, b1, b2)

	var a2 int64 = 9999
	var b21 int8 = int8(a2)
	fmt.Printf("a2=%v b21=%v\n", a2, b21)
	fmt.Printf("a2=%T b21=%T\n", a2, b1)//溢出处理,转换时需要考虑范围


	/* ************** 
	测试 题1 -- 不能编译通过
	var n1 int32 = 12
	var n2 int64
	var n3 int8
	n2 = n1 + 20//应该这样处理 int64(n1 + 20) 或者 n2 = int64(n1) + 20
	n3 = n1 + 20//应该这样处理 int8(n1 + 20) 或者 n3 = int8(n1) + 20
	*/

	/* ************** 
	测试 题2 -- 不能编译通过
	var n1 int32 = 12
	var n3 int8
	var n4 int8

	n4 = int8(n1) + 127//【编译通过，但是结果不是 127+12 会溢出】
	n3 = int8(n1) + 128//【编译不通过，因为128超出int8(-128~127)】

	fmt.Printf("n3=%v n4=%v", n3, n4)
	*/
	
	

}
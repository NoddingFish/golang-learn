package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//TODO 算术运算符
	var num int
	var num2 int = 64
	var num3 int = 31
	var flo1 float32
	// num = num2 - num3
	num = num2 % num3
	num++
	num++
	num--

	fmt.Println("======================= 算术运算符 =====================")
	fmt.Println("结果：", num)

	fmt.Println("（整数）结果：", 10 / 4)
	fmt.Println("（舍弃小数）结果：", 19 / 5) // 结果是3.8 但是实际结果是 3

	flo1 = 10 / 4
	fmt.Println("（整数）结果2：", flo1)

	var flo2 float32 = 10.0 / 4
	fmt.Println("（小数）结果1：", flo2)

	//%的使用： 公式： a % b = a - a / b * b
	fmt.Println("（取余）结果1：", 10 % 3)
	fmt.Println("（取余）结果2：", -10 % 3)
	fmt.Println("（取余）结果3：", 10 % -3)
	fmt.Println("（取余）结果4：", -10 % -3)

	//在 golang 中没有 ++a  和 --a 只有 a++ 和 a--

	//课堂练习1：
	var contDay int = 97
	week := contDay / 7
	day := contDay % 7
	fmt.Printf("练习1：97天是：%d个星期%d天 \n", week, day)

	//课堂练习1：
	var huashi float32 = 134.2
	var wendu float32
	wendu = 5.0 / 9 * (huashi - 100.0)

	fmt.Printf("练习2：%v 华氏摄氏度是 %v 摄氏度 \n", huashi, wendu)
}
package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//函数调用的机制
	var num1 float64 = 85
	fmt.Println(jia(num1))
	fmt.Println(num1)

	fmt.Println("====================== 函数多返回值 ======================")
	// res1, _ := sShu(10, 8) // 忽略一个值，可以使用 _ 占位
	res1, res2 := sShu(10.5, 8.2)
	fmt.Printf("两数之和：%v 两数之差是：%.2f \n", res1, res2)

	fmt.Println("====================== 函数递归调用1 ======================")
	diGui(4)

	fmt.Println("====================== 函数递归调用2 ======================")
	diGui2(4)

	fmt.Println("====================== 斐波那契数 ======================")
	fmt.Println("请输入一个斐波那契数：")
	var feiNum int
	fmt.Scanln(&feiNum)
	res := feiBoNaQi(feiNum)
	fmt.Println(res)

	fmt.Println("====================== 解数学函数 ======================")
	math2 := mathFc(5)
	fmt.Println(math2)

	fmt.Println("====================== 猴子吃桃子 ======================")
	fmt.Println(mkEat(1))

	fmt.Println("======================== 取地址 ========================")
	var num01 int = 9
	fmt.Println(test01(&num01))
	fmt.Println(num01)
}

func jia(num1 float64) float64 {
	return num1 + 1
}

func sShu(num1 float64, num2 float64) (float64, float64) {
	res1 := num1 + num2
	res2 := num1 - num2
	return res1, res2
}

func diGui(num int)  {
	if num > 2 {
		num--
		diGui(num)
	}
	fmt.Println("num1=", num)
}

func diGui2(num int)  {
	if num > 2 {
		num--
		diGui2(num)
	} else {
		fmt.Println("num2=", num)
	}
}

func feiBoNaQi(num int) int {
	if num == 1 || num == 2 {
		return 1
	} else {
		return feiBoNaQi(num - 1) + feiBoNaQi(num - 2)
	}
}

func mathFc(num int) int {
	if num == 1 {
		return 3
	} else {
		return 2 * mathFc(num - 1) + 1
	}
}

func mkEat(num int) int {
	if num > 10 || num < 1 {
		fmt.Println("输入的天数不对")
		return 0
	}
	if num == 10 {
		return 1
	} else {
		return (mkEat(num + 1) + 1) * 2
	}
}

func test01(num *int) int {
	*num = *num + 100
	fmt.Println("test01 num = ", *num)
	return *num
}
package main

import (
	"fmt"
)

func main() {
	fmt.Println("============================ 细节1 ============================")
	//TODO 1、切片初始化时 var slice = arr[startIndex: endIndex] 从数组下标为startIndex取到下标为endIndex的元素（不包含 arr[endIndex]）

	//TODO 2、切片也不能越界，但是动态增长之后，是可以使用的

	//TODO 3、切片的简写
	//TODO 3.1、var slice = arr[0:end] => var slice = arr[:end]
	//TODO 3.2、var slice = arr[start:len(arr)] => var slice = arr[start:]
	//TODO 3.3、var slice = arr[0:len(arr)] => var slice = arr[:]
	arr := [...]int{1,2,3,4,5}
	var slice01 []int = arr[1:]
	// slice01 = append(slice01, 9)//TODO! 增加切片
	fmt.Println("slice01=", slice01)

	//TODO 4、cap 内置函数，用于获取切片的容量

	//TODO 5、切片还可以继续切片
	slice02 := slice01[1:2]
	slice02[0] = 99
	fmt.Println("slice02=", slice02)
	fmt.Println("slice01修改后的=", slice01)//! slice01也会改变
	fmt.Println("arr修改后的=", arr)//! arr也会改变

	fmt.Println("============================ 细节2 ============================")
	//TODO append
	slice03 := append(slice01, 9, 10)
	fmt.Println("slice01=", slice01)
	fmt.Println("slice03=", slice03)

	//TODO 通过append将切片slice01追加到slice01
	slice04 := append(slice01, slice01...)//! 后面必须是切片，不能是数组， ... 必须
	fmt.Println("slice04=", slice04)
	fmt.Printf("slice04=%T \n", slice04)

	fmt.Println("============================ 细节3 ============================")
	//TODO 切片的拷贝 copy
	var slice05 = make([]int, 10)
	var slice06 = make([]int, 2)
	copy(slice05, slice04)//! 必须都是切片，数据空间是独立的
	copy(slice06, slice04)//! 不会报错，只会copy slice06长度的切片
	slice04[0] = 20
	fmt.Println("slice04=", slice04)
	fmt.Println("slice05=", slice05)
	fmt.Println("slice06=", slice06)

}
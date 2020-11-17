package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//值类型：int  float  bool string array数组 和 struct结构体
	//TODO 值类型数据通常是在栈区
	
	//引用类型：指针 slice切片 map chan管道 interface接口
	//TODO 引用类型通常是在堆区分配空间


	fmt.Println("值类型和引用类型")
}
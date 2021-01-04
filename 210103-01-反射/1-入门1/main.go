package main

import (
	"fmt"
	"reflect"
)

func reflectTest01(inter interface{})  {
	//1、先获取到 reflect.Type
	rTyp := reflect.TypeOf(inter)
	fmt.Println("rTyp =", rTyp) //TODO 本质不是 int 类型，可以调用方法

	//2、获取到 reflect.Value 类型
	rVal := reflect.ValueOf(inter)
	fmt.Println("rVal =", rVal)//TODO 不是真正的100
	fmt.Printf("rVal=%v rVal的类型是：%T \n", rVal, rVal)
	// n2 := 2 + rVal//!报错 mismatched types int and reflect.Value
	// fmt.Println("n2 =", n2)

	//TODO 获取值
	n2 := 2 + rVal.Int()
	fmt.Println("n2 =", n2)

	//将 rVal 转 interface{}
	iV := rVal.Interface()

	//将 interface{} 通过断言转真正的需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

func main()  {
	var num1 int = 100
	reflectTest01(num1)
}
package main

import (
	"fmt"
	"reflect"
)

func reflect01(v interface{})  {
	rVal := reflect.ValueOf(v)
	//rVal 是指针
	fmt.Printf("rVal kind=%v \n", rVal.Kind())

	// rVal.SetInt(20)//! panic: reflect: reflect.Value.SetInt using unaddressable value

	rVal.Elem().SetInt(20)
}

func main()  {
	var num int = 10
	reflect01(&num)
	fmt.Println("num=", num)
}
package main

import (
	"fmt"
	"reflect"
)

func reflectTest02(v interface{})  {
	rTyp := reflect.TypeOf(v)
	fmt.Println("rTyp = ", rTyp)

	rVal := reflect.ValueOf(v)
	fmt.Println("rVal = ", rVal)
	fmt.Printf("rVal 的类型是%T\n", rVal)

	//TODO 细节：获取 reflect.Type 和 reflect.Value 对应的 kind
	kind1 := rTyp.Kind() // 常量
	kind2 := rVal.Kind() // 常量
	fmt.Println("kind1 = ", kind1)
	fmt.Println("kind2 = ", kind2)

	iV := rVal.Interface()
	fmt.Printf("iV=%v type=%T\n ", iV, iV)//TODO 反射是运行时
	// fmt.Printf("iV=%v type=%T name=%v\n ", iV, iV, iV.Name)//! 报错

	stu, ok := iV.(Student)//类型断言
	if ok {
		fmt.Printf("stu=%v type=%T name=%v\n ", stu, stu, stu.Name)
	}
}

type Student struct {
	Name string
	Age int
}

func main()  {
	var stu Student = Student{
		Name: "tom~",
		Age: 28,
	}
	reflectTest02(stu)
}
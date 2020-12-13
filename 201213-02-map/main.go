package main

import (
	"fmt"
)

func main() {
	//TODO map映射 key->value的集合
	//TODO key 的类型可以是：bool 数字 string 指针 channel 接口 结构体 数组
	//TODO! key 的类型不能是 slice map function 因为没法用 == 判断
	
	//声明 -不会分配内存的，初始化需要 make ，分配内存后才可以使用
	// var a map[string]string
	// var a map[string]int
	// var a map[int]string
	// var a map[string]map[string]string
	fmt.Println("========================== 方式一 ========================")
	var map1 map[string]string
	map1 = make(map[string]string, 10)//TODO 必须先make ,分配空间
	map1["name"] = "宋江"//!直接使用是会报错的panic: assignment to entry in nil map
	map1["age"] = "35"
	map1["name"] = "武松"//会覆盖
	map1["name2"] = "吴用2"
	map1["name1"] = "吴用1"
	//排序是无序的
	fmt.Println("map1=", map1)

	fmt.Println("========================== 方式二 ========================")
	map2 := make(map[string]int)
	map2["a"] = 1
	map2["b"] = 2
	fmt.Println("map2=", map2)

	
	fmt.Println("========================== 方式三 ========================")
	map3 := map[string]string{
		"hero1": "宋江",
		"hero2": "武松",
	}
	map3["hero3"] = "林冲"
	fmt.Println("map3=", map3)
	
}
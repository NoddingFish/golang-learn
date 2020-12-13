package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("============================= map排序 =============================")
	map1 := make(map[int]int)
	map1[10] = 10
	map1[1] = 100
	map1[8] = 800
	map1[4] = 400

	fmt.Println("map1=", map1)

	var keys []int
	for k, _ := range map1 {
		keys = append(keys, k)
	}

	fmt.Println("keys排序前=", keys)
	sort.Ints(keys)
	fmt.Println("keys排序后=", keys)

	for _, k := range keys {
		fmt.Printf("map1[%v]=%v \n", k, map1[k])
	}
}
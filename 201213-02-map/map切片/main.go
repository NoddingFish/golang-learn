package main

import (
	"fmt"
)

func main()  {
	fmt.Println("============================= map切片 =============================")
	var mapSlice1 []map[string]string
	mapSlice1 = make([]map[string]string, 2)
	if mapSlice1[0] == nil {
		mapSlice1[0] = make(map[string]string, 2)
		mapSlice1[0]["name"] = "牛魔王"
		mapSlice1[0]["age"] = "500"
	}

	if mapSlice1[1] == nil {
		mapSlice1[1] = make(map[string]string, 2)
		mapSlice1[1]["name"] = "玉兔精"
		mapSlice1[1]["age"] = "200"
	}

	//TODO 动态增长1
	mapSlice1 = append(mapSlice1, make(map[string]string, 2))
	mapSlice1[2]["name"] = "狐狸精"
	mapSlice1[2]["age"] = "280"

	//TODO 动态增长2
	newMapSlice1 := map[string]string{
		"name": "火云邪神",
		"age": "80",
	}
	mapSlice1 = append(mapSlice1, newMapSlice1)

	fmt.Println(mapSlice1)
}
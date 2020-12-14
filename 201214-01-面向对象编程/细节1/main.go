package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp, rightDown Point
}

type Rect2 struct {
	leftUp, rightDown *Point
}

func main()  {
	//TODO 1、结构体的所有字段在内存中是连续的
	fmt.Println("========================== 结构体的所有字段在内存中是连续的")
	r1 := Rect{Point{1, 2}, Point{3, 4}}
	fmt.Printf("r1.leftUp.x 的地址=%p \n", &r1.leftUp.x)
	fmt.Printf("r1.leftUp.y 的地址=%p \n", &r1.leftUp.y)
	fmt.Printf("r1.rightDown.x 的地址=%p \n", &r1.rightDown.x)
	fmt.Printf("r1.rightDown.y 的地址=%p \n", &r1.rightDown.y)

	fmt.Println()
	
	r2 := Rect2{&Point{1, 2}, &Point{3, 4}}
	fmt.Printf("r2.leftUp 的地址=%p \n", &r2.leftUp)
	fmt.Printf("r2.rightDown 的地址=%p \n", &r2.rightDown)
	//! r2.leftUp r2.rightDown 本身的地址不一定是连续的
}
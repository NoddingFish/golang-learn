package main

import (
	"fmt"
)

type Box struct {
	len float64
	width float64
	height float64
}

func (b Box) getVol() float64  {
	return b.len * b.width * b.height
}

func main()  {
	//TODO 工厂模式
	var box1 Box
	box1.len = 1.1
	box1.width = 2.0
	box1.height = 4.0
	vol1 := box1.getVol()
	fmt.Println("体积为：", vol1)

	//方式二
	box2 := Box{
		width:2.0,
		len: 1.1,
		height:4.0,
	}
	vol2 := box2.getVol()
	fmt.Println("体积为：", vol2)

	//方式三
	box3 := Box{1.1, 2, 4.0}
	vol3 := box3.getVol()
	fmt.Println("体积为：", vol3)

	//方式四
	box4 := &Box{1.1, 2, 4.0}
	// vol4 := (*box4).getVol() 相同
	vol4 := box4.getVol()
	fmt.Println("体积为：", vol4)
	fmt.Println("box4为：", box4)
	fmt.Println("*box4为：", *box4)

	box5 := &Box{
		width:2.0,
		len: 1.1,
		height:4.0,
	}
	vol5 := box5.getVol()
	fmt.Println("体积为：", vol5)
	fmt.Println("box5为：", box5)
	fmt.Println("*box5为：", *box5)
}
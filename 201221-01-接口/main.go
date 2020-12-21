package main

import (
	"fmt"
)

//声明一个接口
type Usb interface {
	//声明了两个没有实现的方法
	start()
	stop()
}

type Usb2 interface {
	//声明了两个没有实现的方法
	start()
	stop()
	// test()//!如果没有 test方法，则会报错
}

type Phone struct {

}
func (phone Phone) start()  {
	fmt.Println("手机开始工作...")
}
func (phone Phone) stop()  {
	fmt.Println("手机停止工作...")
}

type Camera struct {

}
func (c Camera) start()  {
	fmt.Println("手机开始工作...")
}
func (c Camera) stop()  {
	fmt.Println("手机停止工作...")
}

type Computer struct {

}
//编写一个方法working方法，接受一个Usb接口类型
func (c Computer) Working(usb Usb2)  {
	usb.start()
	usb.stop()
}


func main()  {
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	computer.Working(phone)
	fmt.Println("====================1=")
	computer.Working(camera)
}
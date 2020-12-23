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
	Name string
}
func (phone Phone) start()  {
	fmt.Println("手机开始工作...")
}
func (phone Phone) stop()  {
	fmt.Println("手机停止工作...")
}
func (phone Phone) Call()  {
	fmt.Println("手机 正在通话中...")
}

type Camera struct {
	Name string
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
	//TODO 如果usb 是指向 Phone 变量的，还需要调用一个 Call 方法
	//需要类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.stop()

}



func main()  {
	//TODO 变量（实例）具有多种形态，多态特征是通过接口实现的
	//TODO 之前的 USB 的案列， working(usb Usb) 变量会根据传入的参数，usb体现出多态的特点

	//TODO 多态数组
	var subArr [3]Usb
	subArr[0] = Phone{Name:"vivo"}
	subArr[1] = Phone{Name:"小米"}
	subArr[2] = Camera{Name:"佳能"}
	fmt.Println("subArr = ", subArr)

	var computer Computer
	for _ , v := range subArr {
		computer.Working(v)
		fmt.Println()
	}
}
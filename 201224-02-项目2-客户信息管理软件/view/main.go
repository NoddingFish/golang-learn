package main

import (
	"fmt"
)

type customerView struct {
	key int
	loop bool//表示是否循环显示菜单
}

func (cv *customerView) mainMenu()  {
	for {
		fmt.Println("----------------------- 客户信息管理软件 -----------------------")
		fmt.Println("                           1 添加客户")
		fmt.Println("                           2 修改客户")
		fmt.Println("                           3 删除客户")
		fmt.Println("                           4 客户列表")
		fmt.Println("                           4 退    出")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&cv.key)

		switch cv.key {
		case 1:
			fmt.Println("1...")
			// cv.showDetails()
		case 2:
			fmt.Println("2...")
			// cv.income()
		case 3:
			fmt.Println("3...")
			// cv.pay()
		case 4:
			fmt.Println("4...")
			// cv.exit()
		case 5:
			fmt.Println("5...")
			// cv.exit()
		default:
			fmt.Println("请输入正确的选项...")
		}

		if !cv.loop {
			break//退出成功
		}
	}
	fmt.Println("你退出了客户信息管理软件...")
}

func main()  {
	customerView := customerView{
		key: 0,
		loop: true,
	}
	customerView.mainMenu()
}
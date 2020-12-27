package main

import (
	"fmt"
	"Learn/201224-02-项目2-客户信息管理软件/service"
	"Learn/201224-02-项目2-客户信息管理软件/model"
)

type customerView struct {
	key int
	loop bool//表示是否循环显示菜单
	customerService *service.CustomerService
}

func (this *customerView) list()  {
	customer := this.customerService.List()
	fmt.Println("------------------------- 客户列表 -----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i := 0; i < len(customer); i++ {
		fmt.Println(customer[i].GetInfo())
	}
	fmt.Println("----------------------- 客户列表完成 -----------------------")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func (this *customerView) add()  {
	fmt.Println("------------------------- 添加客户 -----------------------")
	fmt.Print("姓名：")
	name := ""
	fmt.Scanln(&name)

	fmt.Print("性别：")
	gender := ""
	fmt.Scanln(&gender)

	fmt.Print("年龄：")
	age := 0
	fmt.Scanln(&age)

	fmt.Print("电话：")
	phone := ""
	fmt.Scanln(&phone)

	fmt.Print("Email：")
	email := ""
	fmt.Scanln(&email)

	//注意id没有输入，需要是唯一的，系统分配
	customers := model.NewCustomer2(name, gender, age, phone, email)
	if this.customerService.Add(customers) {
		fmt.Println("----------------------- 添加完成 -----------------------")
	} else {
		fmt.Println("----------------------- 添加失败 -----------------------")
	}
	fmt.Println()
	fmt.Println()
	fmt.Println()
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
			cv.add()
		case 2:
			fmt.Println("2...")
			// cv.income()
		case 3:
			fmt.Println("3...")
			// cv.pay()
		case 4:
			fmt.Println("4...")
			cv.list()
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
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
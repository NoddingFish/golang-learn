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

func (this *customerView) delete()  {
	fmt.Println("------------------------- 删除客户 -----------------------")
	fmt.Println("请输入需要删除的客户编号（-1）退出：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return
	}

	fmt.Println("确认是否删除(Y/N)：")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" || choice == "Y" {
		//调用 customerService 的 Delete 方法
		if this.customerService.Delete(id) {
			fmt.Println("------------------------- 删除成功 -----------------------")
		} else {
			fmt.Println("------------------------- 删除失败，请重新输入...")
		}
	}
	
	fmt.Println("----------------------- 客户列表完成 -----------------------")
	fmt.Println()
	fmt.Println()
	fmt.Println()
}

func (this *customerView) exit()  {
	fmt.Println("确认是否退出(Y/N)：")
	choiceKey := ""
	for {
		fmt.Scanln(&choiceKey)
		if choiceKey == "Y" || choiceKey == "y" || choiceKey == "N" || choiceKey == "n" {
			break
		}
		fmt.Println("你的输入有误，确认是否退出(Y/N)：")
	}
	
	if choiceKey == "Y" || choiceKey == "y" {
		this.loop = false
	}
}

func (cv *customerView) mainMenu()  {
	for {
		fmt.Println("----------------------- 客户信息管理软件 -----------------------")
		fmt.Println("                           1 添加客户")
		fmt.Println("                           2 修改客户")
		fmt.Println("                           3 删除客户")
		fmt.Println("                           4 客户列表")
		fmt.Println("                           5 退    出")
		fmt.Print("请选择（1-5）：")

		fmt.Scanln(&cv.key)

		switch cv.key {
		case 1:
			cv.add()
		case 2:
			fmt.Println("2...")
			// cv.income()
		case 3:
			cv.delete()
		case 4:
			fmt.Println("4...")
			cv.list()
		case 5:
			cv.exit()
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
package main

import (
	"fmt"
)

func main()  {
	var key int
	loop := true
	//账户余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//收支的说明
	note := ""
	//收支的详情
	details := "收  支\t账户金额\t收支金额\t说  明"
	flag := false
	for {
		fmt.Println("----------------------- 家庭收支记账软件 -----------------------")
		fmt.Println("                           1 收支明细")
		fmt.Println("                           2 登记收入")
		fmt.Println("                           3 登记支出")
		fmt.Println("                           4 退出软件")
		fmt.Print("请选择（1-4）：")

		fmt.Scanln(&key)

		switch key {
		case 1:
			if flag {
				fmt.Println("----------------------- 当前收支明细记录 -----------------------")
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细记录，去记录一笔吧...")
			}
			
		case 2:
			fmt.Println("登记收入：")
			fmt.Print("本次收入金额：")
			fmt.Scanln(&money)
			balance += money//修改账户余额
			fmt.Print("本次收入说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true
		case 3:
			fmt.Println("登记支出：")
			fmt.Print("本次支出金额：")
			fmt.Scanln(&money)
			if money > balance {
				fmt.Println("余额不足...请重新输入...")
				break
			}
			balance -= money//修改账户余额
			fmt.Print("本次支出说明：")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
			flag = true
		case 4:
			fmt.Print("确定要退出家庭收支记账软件吗？ y/n ")
			isExit := ""
			for {
				fmt.Scanln(&isExit)
				if isExit == "y" || isExit == "n" {
					break
				}
				fmt.Print("输入有误，确定要退出家庭收支记账软件吗？ y/n ")
			}
			if isExit == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项...")
		}

		if !loop {
			break//退出成功
		}
	}

	fmt.Println("你已经退出了家庭收支记账软件...")
}
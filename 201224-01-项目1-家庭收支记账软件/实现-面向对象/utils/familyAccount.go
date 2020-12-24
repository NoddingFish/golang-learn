package utils

import (
	"fmt"
)

type FamilyAccount struct {
	key int
	loop bool
	//账户余额
	balance float64 
	//每次收支的金额
	money float64 
	//收支的说明
	note string
	//收支的详情
	details string
	flag bool
}

func (this *FamilyAccount) MainMenu()  {
	for {
		fmt.Println("----------------------- 家庭收支记账软件 -----------------------")
		fmt.Println("                           1 收支明细")
		fmt.Println("                           2 登记收入")
		fmt.Println("                           3 登记支出")
		fmt.Println("                           4 退出软件")
		fmt.Print("请选择（1-4）：")

		fmt.Scanln(&this.key)

		switch this.key {
		case 1:
			this.showDetails()
		case 2:
			this.income()
		case 3:
			this.pay()
		case 4:
			this.exit()
		default:
			fmt.Println("请输入正确的选项...")
		}

		if !this.loop {
			break//退出成功
		}
	}
	
	fmt.Println("你已经退出了家庭收支记账软件...")
}

func (this *FamilyAccount) showDetails()  {
	if this.flag {
		fmt.Println("----------------------- 当前收支明细记录 -----------------------")
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细记录，去记录一笔吧...")
	}
}

func (this *FamilyAccount) income()  {
	fmt.Println("登记收入：")
	fmt.Print("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money//修改账户余额
	fmt.Print("本次收入说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) pay()  {
	fmt.Println("登记支出：")
	fmt.Print("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足...请重新输入...")
		// break
	}
	this.balance -= this.money//修改账户余额
	fmt.Print("本次支出说明：")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

func (this *FamilyAccount) exit()  {
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
		this.loop = false
	}
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key: 0,
		loop: true,
		balance: 10000.0,
		money: 0.0,
		note: "",
		flag: false,
		details: "收  支\t账户金额\t收支金额\t说  明",
	}
}
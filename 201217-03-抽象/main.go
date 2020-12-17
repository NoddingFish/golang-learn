package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd string
	Balance float64
}

//方法
//1、存款
func (a *Account) Add(money float64, pwd string)  {
	if pwd != a.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	a.Balance += money
	fmt.Println("存款成功~")
}

//2、取款
func (a *Account) Get(money float64, pwd string)  {
	if pwd != a.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	if money <= 0 {
		fmt.Println("你输入的金额不正确")
		return
	}
	a.Balance -= money
	fmt.Println("取款成功~")
}

//3、查询余额
func (a *Account) Query(pwd string)  {
	if pwd != a.Pwd {
		fmt.Println("你输入的密码不正确")
		return
	}
	fmt.Println("余额是：", a.Balance)
}

func main()  {
	
	account := Account{
		AccountNo: "GS10001",
		Pwd: "666666",
		Balance: 100.1,
	}
	pwd := "666666"

	fmt.Println("account是：", account)

	account.Add(20, pwd)
	account.Query(pwd)
	account.Get(28, pwd)
	account.Query(pwd)
}
package service

import (
	"Learn/201224-02-项目2-客户信息管理软件/model"
)

//该 customerService 需要完成对 Customer的操作，增删改查
type CustomerService struct {
	customers []model.Customer
	customerNum int//表示该切片有多少客户
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112", "zhangsan@sohu.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

//返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

//添加客户倒 customer 中
func (this *CustomerService) Add(customer model.Customer) bool {
	//确定一个分配id的规则
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}
package service

import (
	"Learn/201224-02-项目2-客户信息管理软件/model"
)

//该 customerService 需要完成对 Customer的操作，增删改查
type CustomerService struct {
	customers []model.Customer
	customerNum int//表示该切片有多少客户
}
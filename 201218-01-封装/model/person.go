package model

import (
	"fmt"
)

type person struct {
	Name string
	age int //其他包不能直接访问...
	sal float64 
}


func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

//访问age 和 sal
func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年龄范围不正确")
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSal(sal float64) {
	if sal > 0 && sal < 300000 {
		p.sal = sal
	} else {
		fmt.Println("薪资范围不正确")
	}
}

func (p *person) GetSal() float64 {
	return p.sal
}
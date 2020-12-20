package main

import (
	"fmt"
	"Learn/201218-01-封装/model"
)

type Person struct {
	Age string
	Pwd string
	Balance float64
}

func main()  {
	p := model.NewPerson("tom")
	p.SetAge(18)
	p.SetSal(30000)
	fmt.Println("person=", p)
	fmt.Println("person.GetAge=", p.GetAge())
	fmt.Println("person.GetSal=", p.GetSal())
}
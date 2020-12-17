package main

import (
	"fmt"
	"Learn/201217-02-工厂模式/model"
)

func main()  {
	//创建一个student实例 Student 是小写就有问题
	var stu = model.Student{
		Name: "tom", 
		Score: 89.9,
	}
	fmt.Println("stu=", stu)

	//TODO 使用 工厂模式解决
	stuFac := model.NewsStudent("tom~", 90.2)

	fmt.Println("stuFac=", stuFac)
	fmt.Println("stuFac=", *stuFac)
	fmt.Println("stuFac.Name=", stuFac.Name)
	// fmt.Println("stuFac.Name=", stuFac.score)//!会报错
	fmt.Println("stuFac.score=", stuFac.GetScore())//! 解决上面的报错
}
package main

import (
	"fmt"
)

func main()  {
	student := make(map[string]map[string]string, 0)
	student["stu1"] = make(map[string]string, 3)
	student["stu1"]["name"] = "tom"
	student["stu1"]["sex"] = "男"
	student["stu1"]["address"] = "北京市"

	student["stu2"] = make(map[string]string, 3)
	student["stu2"]["name"] = "mary"
	student["stu2"]["sex"] = "女"
	student["stu2"]["address"] = "上海市"

	fmt.Println("student=", student)
	fmt.Println("student[\"stu1\"]=", student["stu1"])
	fmt.Println("student[\"stu1\"][\"address\"]=", student["stu1"]["address"])
}
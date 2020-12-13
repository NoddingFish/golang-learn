package main

import(
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

	student["stu3"] = make(map[string]string, 3)
	student["stu3"]["name"] = "test"
	student["stu3"]["sex"] = "女"


	
	fmt.Println("================================== 查找 ===================================")
	fmt.Println("共有多少对：", len(student))
	fmt.Println("共有多少对：", len(student["stu1"]))
	val, ok := student["stu1"]
	if ok {
		fmt.Println("有这个key，对应的值是：", val)
	} else {
		fmt.Println("没有这个key：", "stu1")
	}

	fmt.Println()
	fmt.Println("================================== 增加 ===================================")
	student["stu3"]["address"] = "杭州市"

	fmt.Println("增加后的：", student)

	fmt.Println()
	fmt.Println("================================== 修改 ===================================")
	student["stu3"]["address"] = "杭州市江干区"

	fmt.Println("修改后的：", student)

	fmt.Println()
	fmt.Println("================================== 删除 ===================================")
	delete(student["stu3"], "address")
	delete(student["stu3"], "sex")
	delete(student["stu3"], "addressaaa")
	fmt.Println("删除后的：", student)

	fmt.Println()
	fmt.Println("================================== 删除所有 ===================================")
	//TODO 1、遍历所有的key，逐一删除
	//TODO 2、直接make一个新的空间
	student = make(map[string]map[string]string, 0)
	fmt.Println("删除所有后的：", student)
}
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


	
	fmt.Println("================================== 遍历 ===================================")
	for key1, value1 := range student {
		fmt.Println("key1=", key1)
		for key2, value2 := range value1 {
			fmt.Printf("\t key2=%v \t value2=%v \n", key2, value2)
		}
	}
}
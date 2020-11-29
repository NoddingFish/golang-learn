package main

import (
	"fmt" //如果我们没有使用一个包，但是不想去掉，前面加 _
)

func main() {
	//控制台输入数据
	//方法一
	var name string
	var age byte
	var sal float32
	var isPass bool
	fmt.Printf("请输入姓名：")
	fmt.Scanln(&name)

	fmt.Printf("请输入年龄：")
	fmt.Scanln(&age)

	fmt.Printf("请输入薪水：")
	fmt.Scanln(&sal)

	fmt.Printf("请输入是否通过考试：")
	fmt.Scanln(&isPass)

	fmt.Printf("姓名:%v 年龄:%v 薪水:%v 是否通过考试:%v \n", name ,age, sal, isPass)

	//方法二
	fmt.Printf("请输入你的姓名，年龄，薪水，是否通过考试，使用空格隔开。")
	fmt.Scanf("%s %d %f %t", &name, &age, &sal, &isPass)
	fmt.Printf("姓名:%v 年龄:%v 薪水:%v 是否通过考试:%v \n", name ,age, sal, isPass)
}
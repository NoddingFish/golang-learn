package testcase1

import (
	"fmt"
	"testing" //轻量的测试框架
)

//TODO 使用：  go test -v 【测试所有，运行正确或错误，都会输出日志】
//TODO 使用：  go test 【运行正确，没有日志，错误，会输出日志】
//TODO 使用：  go test -v cat_test.go cat.go 【测试单个文件】
//TODO 使用：  go test -v -run TestAddUpper 【测试单个方法】

//TODO Test后的第一个字母需要大写
func TestAddUpper(t *testing.T)  {
	res := addUpper(10)
	if res != 55 {
		//错误 日志
		t.Fatalf("addUpper(10) 执行错误，期望值是：%v 实际值是：%v", 55, res)
	}

	t.Logf("addUpper(10) 执行正确...")
}

func TestHello(t *testing.T)  {
	fmt.Println("TestHello被调用")
}
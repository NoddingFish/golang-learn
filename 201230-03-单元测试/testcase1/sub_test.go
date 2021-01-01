package testcase1

import (
	_ "fmt"
	"testing" //轻量的测试框架
)


func TestGetSub(t *testing.T)  {
	res := getSub(10, 3)
	if res != 7 {
		//错误 日志
		t.Fatalf("getSub(10, 3) 执行错误，期望值是：%v 实际值是：%v", 55, res)
	}

	t.Logf("getSub(10, 3) 执行正确...")
}

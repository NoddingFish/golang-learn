package case2

import (
	"testing"
)

func TestStore(t *testing.T) {
	monster := &Monster{
	// monster := Monster{ // 不写 & 也可以
		Name: "红孩儿",
		Age: 10,
		Skill: "吐火",
	}

	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store() 错误")
	}
	t.Logf("monster.Store() 测试成功")
}

func TestReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() 错误")
	}
	//进一步判断
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() 错误")
	}
	t.Logf("monster.ReStore() 测试成功")
}
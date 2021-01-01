package case2

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Skill string
}

func (mon *Monster) Store() bool {
	// monster := Monster{
	// 	Name: "牛魔王",
	// 	Age: 500,
	// 	Skill: "牛魔拳",
	// }

	res, err := json.Marshal(mon)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return false
	}
	// fmt.Println(string(res))
	//保存文件
	filePath := "./file/monster.log"
	err = ioutil.WriteFile(filePath, res, 0666)
	if err != nil {
		fmt.Println("write file err:", err)
		return false
	}
	return true
}

func (mon *Monster) ReStore() bool {
	//先从文件中读取 json
	filePath := "./file/monster.log"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err:", err)
		return false
	}
	//使用读取到的 []byte，反序列化
	err = json.Unmarshal(data, mon)
	if err != nil {
		fmt.Println("Unmarshal err:", err)
		return false
	}
	return true
}
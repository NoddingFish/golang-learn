package main

import (
	"fmt"
)

func modifyUser(users map[string]map[string]string, name string)  {
	// v, ok := users[name]
	if users[name] != nil {
		users[name]["pwd"] = "888888"
	} else {
		//没有用户
		users[name] = make(map[string]string, 0)
		users[name]["pwd"] = "888888"
		users[name]["nick"] = name
	}
}

func main()  {
	users := make(map[string]map[string]string, 0)
	users["smith"] = make(map[string]string, 0)
	users["smith"]["pwd"] = "000000"
	users["smith"]["nick"] = "小胡毛"
	modifyUser(users, "tom")
	modifyUser(users, "mary")
	modifyUser(users, "smith")
	fmt.Println("users=", users)
}
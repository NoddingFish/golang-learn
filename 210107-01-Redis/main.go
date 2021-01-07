package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	//1、链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()//TODO 非常重要，关闭
	fmt.Println("redis 连接成功...", conn)

	//*2、通过go向redis写入数据 == string
	fmt.Println("==================== redis string")
	_, err = conn.Do("set", "name", "jack01")
	if err != nil {
		fmt.Println("set err=", err)
		return
	}
	fmt.Println("redis set 成功")

	//TODO conn.Do("get", "name") 返回的是 interface{}
	//TODO redis.String() 提供自转的方法
	name, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("get err=", err)
		return
	}
	fmt.Println("redis get 成功 name = ", name)

	//*3、通过go向redis写入数据 == hash
	fmt.Println("==================== redis hash")
	_, err = conn.Do("hSet", "user01", "name", "tom 你好!")
	if err != nil {
		fmt.Println("hSet err=", err)
		return
	}
	fmt.Println("redis hSet 成功")

	//TODO conn.Do("get", "name") 返回的是 interface{}
	//TODO redis.String() 提供自转的方法
	user01Name, err := redis.String(conn.Do("hGet", "user01", "name"))
	if err != nil {
		fmt.Println("hGet err=", err)
		return
	}

	//TODO redis.Strings() //返回集合
	user01, err := redis.Strings(conn.Do("hGetAll", "user01"))
	if err != nil {
		fmt.Println("hGet err=", err)
		return
	}

	fmt.Println("redis hGet 成功 user01Name = ", user01Name)
	fmt.Println("redis hGet 成功 user01 = ", user01)
}
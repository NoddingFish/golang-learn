package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

type Person struct {
	Name string
	Age int
	Address string
}

func create() {
	personChan := make(chan Person, 10)
	
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		person := Person{
			Name: "测试" + strconv.Itoa(rand.Intn(100)),
			Age: 19 + i,
			Address: "浙江省杭州市江干区" + strconv.Itoa(i+1) + "幢",
		}

		fmt.Println(person)

		personChan<- person
	}

	fmt.Println(personChan)

	personChanNum := len(personChan)
	for i := 0; i < personChanNum; i++ {
		fmt.Println("第", i, "个是：", <-personChan)
	}
}

func main()  {
	create()
}
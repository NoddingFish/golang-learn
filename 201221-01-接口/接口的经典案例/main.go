package main

import (
	"fmt"
	"sort"
	"math/rand"
)
//1、声明 hero 结构体
type Hero struct {
	Name string
	Age int
}

//2、声明一个 Hero 结构体的切片类型
type HeroSlice []Hero

//3、实现 interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}
//Less() 决定使用什么标准排序
//1、按照 hero 的年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	return hs[i].Age < hs[j].Age
	//比如说：现在需要按照姓名排序 -- 只需要修改这里
	// return hs[i].Name < hs[j].Name
}

//Less() 决定使用什么标准排序
func (hs HeroSlice) Swap(i, j int) {
	//交换数据
	// temp := hs[i]
	// hs[i] = hs[j]
	// hs[j] = temp

	//TODO 等价于
	hs[i], hs[j] = hs[j], hs[i]
}

func main()  {
	//TODO 实现对 hero 结构体切片的排序： sort.Sort(data interface)
	var intSlice = []int{0, -2, 10, 89, 31, 20}
	//冒泡排序
	//系统排序
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//要求：对结构体切片进行排序
	var heros HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄~%d", rand.Intn(100)),
			Age: rand.Intn(100),
		}
		//将 hero append 到 heroSlice 切片
		heros = append(heros, hero)
	}
	fmt.Println("排序前 heros = ", heros)
	//调用接口排序
	sort.Sort(heros)
	fmt.Println("排序后 heros = ", heros)


	fmt.Println("======================= 交换 =======================")
	i := 10
	j := 20
	i , j  = j ,i
	fmt.Println("i j = ", i, j)
}
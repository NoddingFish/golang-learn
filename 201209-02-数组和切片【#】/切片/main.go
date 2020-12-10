package main

import(
	"fmt"
)

func main(){
	//TODO 切片（slice）：可以理解是动态的数组，个数不确定
	//TODO 切片是数组的一个引用，是引用类型，引用传递
	//TODO 切片和数组类似
	//TODO 切片的长度可以变化的
	//TODO 切片的基本语法 var slice01 []int

	//TODO slice 从底层来说，其实就是一个数据结构体（struct 结构体）
	//TODO type slice struct {
	//TODO 		ptr *[2]int //指针的地址
	//TODO 		len int     //长度
	//TODO 		cap int		//容量
	//TODO }

	fmt.Println("==================================== 基本使用 1 ====================================")
	var arr01 [5]int = [...]int{11, 22, 33, 44, 55}
	slice01 := arr01[1:3]
	//上面代码表示 slice01 引用到 arr01 这个数组下标1和3（但是不包含3）
	fmt.Println("arr01=", arr01)
	fmt.Println("slice01=", slice01)
	fmt.Println("slice01 len=", len(slice01))
	fmt.Println("slice01 容量=", cap(slice01))//切片的容量可以动态变化

	slice01[1] = 34
	fmt.Println("arr01也会修改=", arr01)
	fmt.Printf("arr01[1]的地址是  ：%p \n", &arr01[1])
	fmt.Printf("slice01[0]的地址是：%p slice01[1]数据是：%v \n", &slice01[0], slice01[1])

	fmt.Println()
	fmt.Println("==================================== 基本使用 2 ====================================")
	//TODO make 创建切片，也会创建一个数组，由切片在底层维护，程序员不可见
	var slice02 []int = make([]int, 4, 10)//参数分别：切片类型，长度，容量（可选）  默认值全是0
	slice02[1] = 100
	fmt.Println("slice02=", slice02)
	fmt.Printf("slice02类型=%T \n", slice02)
	fmt.Println("slice02长度=", len(slice02))
	fmt.Println("slice02容量=", cap(slice02))

	fmt.Println()
	fmt.Println("==================================== 基本使用 3 ====================================")
	var slice03 []string = []string{"tom", "jack", "john"}
	fmt.Println("slice03=", slice03)
	fmt.Printf("slice03类型=%T \n", slice03)
	fmt.Println("slice03长度=", len(slice03))
	fmt.Println("slice03容量=", cap(slice03))

	fmt.Println()
	fmt.Println("==================================== 切片的遍历 ====================================")
	var slice04 []int = []int{11, 22, 33, 44, 55}
	for i := 0; i < len(slice04); i++ {
		fmt.Println("切片遍历1:", slice04[i])
	}
	for index, value := range slice04 {
		fmt.Println("切片遍历2:", index, value)
	}
}
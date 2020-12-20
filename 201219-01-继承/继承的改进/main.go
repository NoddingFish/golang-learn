package main

import (
	"fmt"
)

//共有的
type Student struct {
	Name string
	age int
	Score float64
}

func (p *Student) showInfo()  {
	fmt.Println("学生的信息：", *p)
}

func (p *Student) SetScore(score float64)  {
	p.Score = score
}

//*小学生
type Pupil struct {
	Student//TODO 嵌入了Student匿名结构体-继承
}
func (p *Pupil) testings()  {
	fmt.Println("小学生正在考试...")
}

//*大学生
type Pupil2 struct {
	Student//TODO 嵌入了Student匿名结构体-继承
	age int
}
func (p *Pupil2) testings()  {
	fmt.Println("大学生正在考试...")
}
//! 细节3
func (p *Pupil2) showInfo()  {
	fmt.Println("== 细节3、大学生的信息：", *p)
}

func main() {
	//TODO 继承的使用
	var pupil = &Pupil{}
	pupil.Student.Name = "tom~"
	pupil.Student.age = 22//TODO 细节1：结构体可以使用继承的匿名结构体内的所有的属性和方法，包括私有的
	pupil.Student.SetScore(91)
	pupil.Student.showInfo()//TODO 细节1：结构体可以使用继承的匿名结构体内的所有的属性和方法，包括私有的
	pupil.testings()
	
	//TODO 细节2、上面的写法 pupil.Student. 可以简写成 pupil.
	var pupil2 = &Pupil2{}
	pupil2.Name = "mary~"
	pupil2.age = 26
	pupil2.SetScore(99)
	pupil2.showInfo()
	pupil2.testings()

	//TODO 细节3、当结构体和继承的匿名结构体有相同的字段或方式是，编译器采用就近访问原则，
	//TODO        如希望访问匿名机构提的字段和方法，可以通过匿名结构体名来区分

	//TODO 细节4、如果一个结构体同时继承了两个匿名结构体，那两个结构体包含同名的属性和方法，而
	//TODO        而结构体本身没有同名的属性和方法，那访问时，必须明确指明匿名结构体，否则报错

}
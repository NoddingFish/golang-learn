package main

import (
	"fmt"
)

//小学生
type Pupil struct {
	Name string
	Age int
	Score float64
}

func (p *Pupil) ShowInfo()  {
	fmt.Println("学生的信息：", *p)
}

func (p *Pupil) SetScore(score float64)  {
	p.Score = score
}

func (p *Pupil) testings()  {
	fmt.Println("小学生正在考试...")
}

//大学生 、 研究生 ...
//! 会出现重复代码，代码冗余
type Pupil2 struct {
	Name string
	Age int
	Score float64
}

func (p *Pupil2) ShowInfo()  {
	fmt.Println("学生的信息：", *p)
}

func (p *Pupil2) SetScore(score float64)  {
	p.Score = score
}

func (p *Pupil2) testings()  {
	fmt.Println("大学生正在考试...")
}

func main() {
	//TODO 继承
	var pupil = &Pupil{
		Name: "tom",
		Age:18,
	}

	pupil.testings()
	pupil.SetScore(78)
	pupil.ShowInfo()

	var pupil2 = &Pupil2{
		Name: "tom2",
		Age:19,
	}
	pupil2.testings()
	pupil2.SetScore(90)
	pupil2.ShowInfo()

}
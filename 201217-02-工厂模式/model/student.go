package model

//定义一个结构体
type Student struct {
	Name string
	Score float64
}

//定义一个结构体 首字母是小写，只能在model包使用
type studentFac struct {
	Name string
	score float64 
}

//TODO 使用工厂模式解决
func NewsStudent(n string, s float64) *studentFac  {
	return &studentFac{
		Name: n,
		score: s,//如果是小写在其他的包不能直接访问
	}
}

func (s *studentFac) GetScore() float64 {
	return s.score
}
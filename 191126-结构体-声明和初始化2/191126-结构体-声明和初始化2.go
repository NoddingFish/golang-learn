package main

import "fmt"


type Saiyans struct {
	Name string
	Power int
}

func main()  {
	goku := &Saiyans{"Power", 9000}
	super(goku)
	fmt.Println(goku.Power)
}

func super(s *Saiyans)  {
	s.Power += 10000
}

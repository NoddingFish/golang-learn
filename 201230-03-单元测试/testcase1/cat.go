package testcase1

import (
	_ "fmt"
)

func addUpper(n int) int  {
	res := 0
	for i := 1; i <= n - 1; i++ {
		res += i
	}
	return res
}

func getSub(n1 int, n2 int) int  {
	return n1 - n2
}

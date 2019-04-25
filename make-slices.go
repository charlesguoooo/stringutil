package stringutil

//切片可以用内建函数make来创建，这也是你创建动态数组的方式

import (
	"fmt"
)

func MS() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

package stringutil

//切片可以用内建函数make来创建，这也是你创建动态数组的方式
//make函数接受元素类型，长度和容量（可选）作为传入参数

import (
	"fmt"
)

func MS() {
	a := make([]byte, 20)
	copy(a,"hello,workd!")
	printSlice("Hello,world!", a)
	//
	//b := make([]int, 0, 5)
	//printSlice("b", b)
}

func printSlice(s string, x []byte) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), string(x))
}

package stringutil

import (
	"fmt"
)

/*
可以在函数中声明一个匿名函数类型的变量，这个匿名函数成为闭包
它可以引用其函数体之外的变量，例如下边的adder函数，每个闭包都被绑定在格子的sum变量上
闭包是由函数及其相关引用环境组合而成的实体(即：闭包=函数+引用环境)

// 匿名函数分为带返回值和不带返回值，带参数和不带参数
// 带返回值
func main() {
	a := func() int {
		x := 0
		return x + 1
	}
	fmt.Println(a())
}

// 带参数
func main() {
	a := func(arg int) int {
		x := 0
		return x + arg
	}
	fmt.Println(a(2))
}

*/
func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum = sum+x
		return sum
	}
}

func FC1() {
	pos,neg := adder(),adder()
	for i := 0;i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
			)
	}
	fmt.Printf("%p\n",pos)
	fmt.Printf("%p\n",neg)
}


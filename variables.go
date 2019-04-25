package stringutil

import (
	"fmt"
)

// = 赋值,因此之前必须声明
// := 声明并赋值
// 函数外的每个语句都必须以(var,func...)关键字开始,:=不能使用在函数外
var a int = 20

func Vab() {
	//a = 30
	a := 10
	b, c := 20, 0
	fmt.Printf("main() a = %d\n", a)
	c = sum(a, b)
	fmt.Printf("main() c = %d\n", c)

}

// 这里a，b的值来自上面的行参
func sum(a, b int) int {
	fmt.Printf("sum() a = %d\n", a)
	fmt.Printf("sum() b = %d\n", b)
	return a + b
}

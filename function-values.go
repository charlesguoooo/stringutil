package stringutil


// 和很多其他语言一样，go中的函数也可以作为值，用作函数的参数或返回值

import (
	"fmt"
	"math"
)

// 比如我们声明一个函数类型，或者说可以直接func compute(fn func(int,int)int) int {}

type FnType func(a float64,b float64) float64

func compute(fn FnType) float64 { // 定义一个函数用来调用一个函数并返回函数调用的结果
	return fn(3,4)
}


func add(int,int) int{
	return 1+1
}
func FV() {
	// 这里定义了一个hypot函数，等会给compute调用，并最终返回调用hypot的结果
	hypot := func(x,y float64)float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5,12))
	fmt.Println(compute(hypot)) //返回hypot(3,4),然后接着往下执行，返回hypot()调用的结果
	fmt.Println(compute(math.Pow)) // fmt.Println(math.Pow(3,4))
}

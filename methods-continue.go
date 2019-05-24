package stringutil

import (
	"fmt"
	"math"
)

//你只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。
//（译注：就是接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。）
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func MC()  {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}

package stringutil

import (
	"fmt"
	"math"
)

// Go 语言提供了另外一种数据类型即接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口
// 接口类型的变量可以保存任何实现了这些方法的值

// 定义接口
type Abser interface {
	Abs() float64
}

// 定义类型
type Vertex6 struct {
	x,y float64
}

// 定义类型
type MyFloat1 float64



// 实现接口方法
func (f *MyFloat1) Abs()  float64 {
	if *f < 0 {
		return float64(-*f)
	}
	return float64(*f)
}

// 实现接口方法
func (v *Vertex6) Abs() float64  {
	return math.Sqrt(v.x*v.x+v.y*v.y)
}

func INT(){
	var a Abser
	f := MyFloat1(-10)
	//v := Vertex5{3,4}

	a = &f
	//a = &v

	//a = v
	fmt.Println(a.Abs())
}



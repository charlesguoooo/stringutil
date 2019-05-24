package stringutil

// 方法即函数，只不过是个带了接收者的函数

import (
	"fmt"
	"math"
)

type Vertex2 struct {
	x,y float64
}

func Abs(v Vertex2) float64  {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func MF() {
	v := Vertex2{3,4}
	fmt.Println(Abs(v))
}


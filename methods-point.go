package stringutil

import (
	"fmt"
	"math"
)

type VERTEX struct {
	X, Y float64
}

func (a VERTEX) abs1() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y)
}

func (a *VERTEX) scale1(f float64) {
	a.X = a.X * f
	a.Y = a.Y * f
}

func METHODSPOINT() {
	a := VERTEX{3, 4}
	a.scale1(10)
	fmt.Println(a.abs1())
}

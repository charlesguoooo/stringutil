package stringutil

import (
	"fmt"
	"math"
)

func st(x float64) float64 {
	z := x
	for i := 0; i < 10; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func ST() {
	fmt.Println(st(2))
	fmt.Println(math.Sqrt(2))
}

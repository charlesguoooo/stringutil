package stringutil

import (
	"fmt"
	"math"
)

func xqrt(x float64) string {
	if x < 0 {
		return xqrt(-x) + "i"
	}
	//fmt.Printf("%T\n",fmt.Sprint(math.Sqrt(x)) )
	return fmt.Sprint(math.Sqrt(x))
}

func Is() {
	fmt.Println(xqrt(2), xqrt(-4))
	//fmt.Printf("%T,%v",fmt.Sprint(10),fmt.Sprint(10))
}

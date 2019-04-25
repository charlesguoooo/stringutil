package stringutil

import (
	"fmt"
	"math"
)

func Tc() {
	var x, y int = 3, 4
	f := math.Sqrt(float64(x*x + y*y))
	// 我们可以试着将int(f)的int删掉再看看z变量的类型
	var z int = int(f)
	fmt.Println(x, y, z)
	fmt.Printf("%T", z)
}

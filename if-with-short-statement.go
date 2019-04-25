package stringutil

import (
	"fmt"
	"math"
)

// 和for一样，if语句可以在条件之前执行短语句
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func IST() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)
}

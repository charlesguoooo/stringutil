package stringutil

import (
	"fmt"
)

const (
	//移位
	// <<从右往左移动10位，空缺的部分填0
	// >>从左往右移9位，空缺部分填0，溢出部分丢弃
	Big   = 1 << 10
	Small = Big >> 9
)

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func Nc() {
	fmt.Printf("%b\n", Big)
	fmt.Printf("%b\n", Small)
	fmt.Printf("%v\n", needInt(Small))

	fmt.Printf("%v\n", needFloat(Small))

	fmt.Printf("%v\n", needFloat(Big))
}

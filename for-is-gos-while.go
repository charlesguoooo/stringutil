package stringutil

import (
	"fmt"
)

// 去掉;，就类似C中的while循环
func Fw() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Print(sum)
}

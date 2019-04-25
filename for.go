package stringutil

import (
	"fmt"
)

func FO() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum = i + sum
	}

	fmt.Print(sum)
}

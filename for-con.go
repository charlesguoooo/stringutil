package stringutil

import "fmt"

func FC() {
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Print(sum)
}

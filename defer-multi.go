package stringutil

import (
	"fmt"
)

func DM() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer println(i)
	}
	fmt.Println("Done")
}

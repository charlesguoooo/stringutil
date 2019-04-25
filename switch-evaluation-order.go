package stringutil

import (
	"fmt"
	"time"
)

// switch 从上到下的顺序执行，当匹配到之后就不往下执行了
func SO() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch monday := time.Monday; monday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

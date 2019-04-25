package stringutil

import (
	"fmt"
	"time"
)

// switch 表达式;表达式其实是可以忽略的，忽略之后相当于switch true

func SC() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening.")
	}
}

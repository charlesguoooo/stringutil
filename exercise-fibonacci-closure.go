package stringutil

import (
	"fmt"
)


func fibonacci() func() int {
	a,b:=0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}

func FIB() {
	f:=fibonacci()
	for i :=0;i<10;i++ {
		fmt.Println(f())
	}
}
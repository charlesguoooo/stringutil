package stringutil

import (
	"fmt"
)

// nil切片的长度和容量都为0，且没有底层数组
func NS() {
	var a []int
	//a := []int{} // 这种方式会有底层数组，因此不会等于nil
	//fmt.Printf("%T\n%T\n",a,s)
	fmt.Println(a, len(a), cap(a))
	if a == nil {
		fmt.Println("nil")
	}
}

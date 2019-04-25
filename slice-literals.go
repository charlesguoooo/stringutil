package stringutil

import (
	"fmt"
)

// 切片文法类似于没有长度的数组文法
func SL() {
	// 这会创建一个数组然后构建了一个基于它的切片
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, false},
		{7, false},
		{11, true},
		{13, true},
	}
	fmt.Println(s)
}

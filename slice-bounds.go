package stringutil

import (
	"fmt"
)

// 在进行切片时，可以利用切片的默认值来忽略上下范围
// 默认最下的范围是0，最上的范围是该切片的长度
func SB() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[1:4]
	fmt.Println(s) // s=357

	s = s[:2]
	fmt.Println(s) // s=35

	s = s[1:]
	fmt.Println(s) // 5
}

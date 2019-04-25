package stringutil

import (
	"fmt"
)

// 切片的长度就是它所包含的元素个数
// 切片的容量是从它的第一个元素开始，到其底层数组元素末尾的个数
func SLC() {
	s := []int{2, 3, 5, 7, 11, 13}
	s = s[2:4] // 它的元素个数是2个[5,7]，容量是4，从[5,7]开始数到数组的最后，即[5,7,11,13]
	fmt.Println(s)
	println(len(s), cap(s))
}

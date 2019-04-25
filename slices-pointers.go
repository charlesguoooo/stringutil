package stringutil

import (
	"fmt"
)

// 切片不存储数据，它只是描述了底层数组中的一段
// 更改切片的元素会彻底修改底层数组中对应的元素
func SP() {
	names := [4]string{
		"john",
		"paul",
		"george",
		"ringo",
	}
	fmt.Printf("%T,%v", names, names)
	a := names[0:2]
	b := names[1:3]
	fmt.Printf("%T,%v,%v", a, a, b)
	b[0] = "x"
	fmt.Println(a, b)
	fmt.Println(names)
}

package stringutil

import "fmt"

// 指定了零个方法的接口值被称为 *空接口：*
// 空接口可以保存任何类型的值
// 空接口被用来处理未知类型的值
// 因为空接口不包含任何方法，所以任何类型都默认实现了空接口

type I3 interface {}
func IE() {
	var i3 I3

	describe3(i3)

	i3 = 42
	describe3(i3)

	i3 = "hello"
	describe3(i3)
}

func describe3(i3 I3) {
	fmt.Printf("(%v,%T)\n",i3,i3)
}
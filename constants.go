package stringutil

import "fmt"

const Pi = 3.14

func Cs() {
	// 常量不能用:=定义，且常量定义后，其值不能再次修改
	const world = "世界"
	fmt.Print("Hello", world)
	fmt.Print("Happy", Pi, "Day")
	const Truth = true
	fmt.Print("Go rules?", Truth)
}

package stringutil

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// %v 相应值的默认格式
// %T 相应值的类型在Go语法中的表示
// %#v 相应值在Go语法中的表示
// %+v 打印结构体时，会添加字段名
// %t 布尔占位符
// %b 二进制表示
// %c 相应Unicode码点所表示的字符
// %d 十进制表示
// %o 八进制表示
// %q 单引号围绕的字符字面值，由Go语法安全转义
// %x 十六进制，字母形式为小写
// %X 十六进制，字母形式为大写
// %U Unicode格式
// %s 字符串占位符
// %f 浮点占位符
func Bt() {
	// const 定义常量，这里将占位符格式定义为常量f
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
}

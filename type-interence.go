package stringutil

import "fmt"

func Ti() {
	// 当我们定义了一个变量不指定类型时，变量类型由右边推断而出
	v := "asdf"
	fmt.Printf("v is of type %T\n", v)
}

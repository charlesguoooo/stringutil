package stringutil

// 类型选择是一种按顺序从几个类型断言中选择分支的结构

// 类型选择与一般的switch语句相似，不过类型选择中的case为类型(而非值),他们根据给定接口值所存储的值的类型进行比较

// 类型选择中的声明与类型断言 i.(T) 的语法相同，只是具体类型 T 被替换成了关键字 type

// this is only used in type switches
import (
	"fmt"
)

func do(i interface{})  {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n",v)
	}
}


func TS(){
	do(21)
	do("hello")
	do(true)
}
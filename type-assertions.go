package stringutil

import "fmt"

// 类型断言提供了访问接口值底层具体值的方式

func TA() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s,ok := i.(string)
	fmt.Println(s,ok)

	f,ok := i.(float64)
	fmt.Println(f,ok)

	f,_ = i.(float64)
	fmt.Println(f)
}
package stringutil

import (
	"fmt"
	"time"
)

// go 使用error值来表示错误状态
// 与fmt.Stringer类似，error类型是一个内建接口


//type error interface {
//	Error() string
//}

type MyError struct {
	When time.Time
	What string
}



func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",e.When,e.What)
}

// 这里run函数返回值的类型是error，这个error类型必须是实现了Error()方法的类型MyError的类型

func run() error {
	// 这里能够直接print出结果是因为print默认做了类型断言
	// 因此如果我们要实现一样的功能，那么我们就需要实现String()方法
	// 详见interfaces-are-satisfied-implicitly
	a := &MyError{time.Now(),"it didn't work",}
	return a
}


func LE() {
	//if err := run();err != nil {
	//	fmt.Println(err)
	//}
	fmt.Println(run())
}

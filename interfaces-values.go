package stringutil

// 接口也是值，它也可以像其他值一样传递

// 接口值可以用作函数参数或返回值

// 变量的类型在声明时指定、且不能改变，称为静态类型。接口类型的静态类型就是接口本身。
// 接口没有静态值，它指向的是动态值。
// 接口类型的变量存的是实现接口的类型的值。该值就是接口的动态值，实现接口的类型就是接口的动态类型

import (
	"fmt"
	"math"
)


type I1 interface {
	M1()
}

type T1 struct {
	S string
}

func (t *T1) M1() {
	fmt.Println(t.S)
}

type F float64

func (f F) M1() {
	fmt.Println(f)
}


func Haha() I1 {
	return &T1{"nihao"}
}


func IV(){
	// i的静态类型是I1
	var i I1

	describe(i)
	// 第一次分配后，i的动态类型是T1,i的动态值是T1类型的值
	i = &T1{"hello"}
	describe(i)
	//i.M1()

	// 第二次分配后，i的动态类型是F，i的动态值是F类型的值
	i = F(math.Pi)
	describe(i)
	//i.M1()
}

func describe(i I1) {
	fmt.Printf("(%v,%T)\n",i,i)
}
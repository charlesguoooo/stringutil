package stringutil

import (
	"fmt"
)

//& 取值地址
//* 指针类型声明或取地址指向的实际值
// Go 没有指针运算
// nil 表示空指针
func PT() {
	i, j := 42, 2701

	p := &i
	//var p *int // 声明指针类型，声明p的类型是[int型的指针]
	//p := new(int) // 这个和上面相等
	//p = &i // 取值地址
	fmt.Printf("%T\n", p)
	fmt.Println(*p) // 取地址指向的实际值

	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}

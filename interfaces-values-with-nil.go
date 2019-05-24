package stringutil

import "fmt"


// 当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil
// 为 nil 接口调用方法会产生运行时错误，因为接口的元组内并未包含能够指明该调用哪个 具体 方法的类型
type I2 interface {
	M2()
}

//type T2 struct {
//	S2 string
//}

//func (t2 *T2) M2() {
//	if t2 == nil {
//		fmt.Println("<Nil>")
//		return
//	}
//	fmt.Println(t2.S2)
//}

func IVWN() {
	var i2 I2
	//var t2 *T2
	//i2 = t2
	describe2(i2)
	//i2.M2()

	//i2 = &T2{"hello"}
	//describe2(i2)
	//i2.M2()
}

func describe2(i2 I2) {
	fmt.Printf("(%v,%T)\n",i2,i2)
}
package stringutil

// 类型通过实现一个接口的所有方法来实现该接口

import (
	"fmt"
)

type I interface {
	M() string
}

type T struct {
	s string
}

func (t *T) M() string {
	return ""
}
func (t *T) String() string {
	return fmt.Sprintf("aa %s",t.s)
}

func run1() I {
	a := &T{"haha"}
	return a
}

func IASI() {
	fmt.Println(run1())
	//i := T{"Hello"}
	//p :=T{"nihao"}
	//fmt.Println(i,p)
	//i.M()
	//p.M()
}
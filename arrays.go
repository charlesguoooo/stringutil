package stringutil

import (
	"fmt"
)

// 数组的长度是数组类型的一部分,[4]int和[5]int是两个数组类型
//看起来像是数组一旦声明了就没法调整，但实际不是，我们可以通过切片的方式来动态体现数组
func ARR() {
	var b [2]int
	b[0] = 1
	b[1] = 2
	a := [2]string{}
	a[0] = "hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Printf("%T\n", a)

	// []type 表示一个切片中的元素的类型为type的切片
	// []int 表示切片中的元素类型为int的切片类型
	// 定义切片和定义数组差不多，唯一不同的就是不需要定义切片长度
	var s []int = primes[1:6]
	s1 := primes[2:5]
	s2 := []string{"c", "b", "a"}
	fmt.Println(s, s1, s2)
	fmt.Printf("%T", s)
}

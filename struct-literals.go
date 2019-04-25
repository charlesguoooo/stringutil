package stringutil

import "fmt"

// 结构体文法表示可以通过结构体字段的值作为列表来分配一个新的结构体
type Vertex1 struct {
	X, Y int
}

// 我们可以通过var来一起操作
var (
	v1 Vertex1  = Vertex1{1, 2}  // 结构体1
	v2          = Vertex1{X: 10} // 结构体2
	v3          = Vertex1{}      //结构体3
	v4 *Vertex1 = &Vertex1{1, 2} //结构体指针4
	v6 *Vertex1 = new(Vertex1)   //结构体指针6
)

func V1() {
	// 也可以在函数里边来用语法糖操作
	v5 := Vertex1{11, 12}
	fmt.Printf("%T,%T,%T,%T,%v", v1, v2, v3, v4, v5)
}

package stringutil

import "fmt"

// 结构体字段用.来访问
type Vertex struct {
	X int
	Y int
}

func (coo *Vertex) GetCoordinate() {
	fmt.Printf("(%d,%d)\n", coo.X, coo.Y)
	return
}

// 这里还有一种方式，叫值拷贝的方式,即当我在通过这个函数操作的时候，他显示的内容只会在这个函数内有效，而用*Vertex则会修改外部的值
func (coo Vertex) GetCoordinate1() {
	fmt.Printf("%d,%d\n", coo.X, coo.Y)
	return
}

func SS() {
	p0 := Vertex{1, 2}
	p0.GetCoordinate()
}

func SS1() {
	p0 := Vertex{Y: 1, X: 2}
	p0.GetCoordinate()
}

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func SS2() {
	// 结构体字段可以通过结构体指针来访问，通过结构体指针的访问是透明的
	p0 := new(Vertex) // p0是指针
	p1 := &Vertex{}   // p1也是个指针
	var t = Vertex{}
	fmt.Printf("%T\n", p1)
	fmt.Printf("%T\n", p0)
	fmt.Printf("%T\n", t)
	fmt.Printf("%T\n", int(1))
	fmt.Println(*p0)
	fmt.Println(*p1)
	p0.X = 1
	p0.Y = 2
	p1.X = 3
	p1.Y = 4
	//p0.GetCoordinate()
	//p1.GetCoordinate()
}

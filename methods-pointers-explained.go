package stringutil

import "fmt"

// 使用指针接收者的原因有二：
//首先，方法能够修改其接收者指向的值。
//其次，这样可以避免在每次调用方法时复制该值。若值的类型为大型结构体时，这样做会更加高效。

// 带指针参数的函数必须接受一个指针
// 而以指针为接收者的方法被调用时，接收者既能为值又能为指针 注1&注2
//type Vertex4 struct {
//	X, Y float64
//}
//
//func Abs1(v Vertex4) float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//// 该函数只能接收指针参数
//func Scale(v *Vertex4, f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func MPE() {
//	v := &Vertex4{3, 4}
//	Scale(v, 10)
//	fmt.Println(Abs1(*v))
//}

//type Vertex4 struct {
//	X, Y float64
//}
//
//func (v *Vertex4) Scale(f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func ScaleFunc(v *Vertex4, f float64) {
//	v.X = v.X * f
//	v.Y = v.Y * f
//}
//
//func MPE() {
//	// 注1
//	v := Vertex4{3, 4}
//	v.Scale(10)
//	//ScaleFunc(&v, 10)
//
//	// 注2
//	p := &Vertex4{4, 3}
//	p.Scale(10)
//	//ScaleFunc(p, 8)
//
//	fmt.Println(v, p)
//}

// 以值为接收者的方法被调用时，接收者既能为值又能为指针 注1&注2
type Vertex4 struct {
	X, Y float64
}

func (v Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func MPE() {
	// 注1
	v := Vertex4{3, 4}
	v.Scale(10)
	//ScaleFunc(&v, 10)

	// 注2
	p := &Vertex4{4, 3}
	p.Scale(10)
	//ScaleFunc(p, 8)

	fmt.Println(v, p)
}

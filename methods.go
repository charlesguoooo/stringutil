package stringutil

// go没有类，不过我们可以为结构体类型定义方法
// 方法的声明和函数类似，他们的区别是：方法在定义的时候，会在func和方法名之间增加一个参数，这个参数就是接收者，这样我们定义的这个方法就和接收者绑定在了一起，称之为这个接收者的方法
// 方法接收者在它自己的参数列表内
import (
	"fmt"
	"math"
)

type MT struct {
	x,y float64
}

// 这是一个函数
func  ok(args float64) float64  {
	return 1.1 + args
}

// 这是一个方法,并且是一个值接收者的方法
func (x MT) ok()  {
	x.x = 12
}

// 这是一个方法，并且是一个指针接收者的方法
func (x *MT) ok1() {
	x.x = 10
}

// func (参数名 参数类型) 方法名  返回值类型
func (v MT) Abs() float64  {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func METHODS() {
	v := MT{3,4}
	//v.ok()
	v.ok1()
	fmt.Println(v.Abs())
}
package stringutil

import (
	"fmt"
	"log"
	"strconv"
)

var t int

func Vab2() {
	// 如果这里用语法糖则表示将t重新声明并赋值了，这个就是局部变量了
	// 如果换成t,err = strconv.Atoi("2")，则表示对声明的全局变量t进行了赋值，因此下面的Vab3就能打出正确的值了
	t, err := strconv.Atoi("2")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("init:", t)
}

func Vab3() {
	fmt.Println("main:", t)
}

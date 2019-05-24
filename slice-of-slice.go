package stringutil

import (
	"fmt"
	"strings"
)
//切片可以包含任何类型，甚至可以包含其他切片
func SOS() {

	a := []string{"_","_","_"} // 返回的是一个包含字符串类型元素的切片
	b := []string{"_","_","_"}
	c := []string{"_","_","_"}
	//board := [][]string{
	//	[]string{"_","_","_"},
	//	[]string{"_","_","_"},
	//	[]string{"_","_","_"},
	//}
	board := [][]string{a,b,c} // 返回的是一个包含切片类型(即上边的那些个包含字符串类型元素的切片)元素的切片

	board[0][0] = "x"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i :=0;i<len(board);i++ {
		fmt.Printf("%s\n",strings.Join(board[i], " "))
	}
}

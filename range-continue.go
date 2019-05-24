package stringutil


import "fmt"

// 可以将下标或值赋予_来忽略
// 若你只需要索引，可以直接忽略第二个变量即可
func RC() {
	pow := make([]int, 10)
	for i,_ := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}

package stringutil

import (
	"fmt"
)

// 向切片中追加元素append(要追加元素的切片,要追加的元素1，2，3)
// 先将旧的slice容量乘以2，如果乘以2后的容量仍小于新的slice容量，则取新的slice容量(append多个elems)
// 如果旧slice容量的2倍大于等于新slice，则取旧slice容量乘以2
// 如果旧的slice容量大于1024，则新slice容量取旧slice容量乘以1.25
func SLICEAPPEND()  {
	var s []int

	printSlice1(s)

	s = append(s,0)
	printSlice1(s)

	s = append(s,1) // old cap 2,2x2>=3
	printSlice1(s)

	s = append(s,2) // new cap 3,因此这里的cap取的是2x2
	printSlice1(s)

}


func printSlice1(s1 []int) {
	fmt.Printf("len=%d cap=%d %v\n",len(s1),cap(s1),s1)
}
package stringutil

import (
	"fmt"
)

// range函数，返回索引和对应的元素
// 如果是字典的话，就返回key和value

var power = []int{1,2,4,8,16,32,64,128}

func RANGE()  {
	for i,v := range power{
		//fmt.Printf("2**%d = %d\n",i,v)
		fmt.Println(i,v)
	}
}
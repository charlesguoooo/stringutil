package stringutil

import (
	"fmt"
	"golang.org/x/tour/reader"
)

type MyReader struct {

}

func (i *MyReader) Read(b []byte) (int,error) {
	//for i := range (b){
	//	b[i] = 'A'
	//}
	//fmt.Println(len(b))
	//return len(b),nil

	// 如果你看了Validate函数，你就会发现只要保证range b[:1]就好了
	// 这个函数里边本身就有循环了
	b[0] = 'A'
	return 1, nil
}

func ER() {
	s := &MyReader{}
	fmt.Println(s.Read([]byte{'A'}))
	reader.Validate(s)
}
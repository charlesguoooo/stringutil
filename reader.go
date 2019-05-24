package stringutil

import (
	"fmt"
	"io"
	"strings"
)

func REA()  {
	r := strings.NewReader("hello, reader!")

	b := make([]byte,8)

// 这是个forever loop
// Read函数会返回int类型和error类型的两个值
// b是byte切片，我们可以通过string(b)来转换成字符串
// io.Reader 表示读取数据流直到结束，当到最后的时候会返回一个io.EOF
	for {
		n,err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n",n,err,b)
		fmt.Printf("b[:n] = %q\n",b[:n])
		if err == io.EOF {
			break
		}
	}
}

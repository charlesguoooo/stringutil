package stringutil

import "fmt"

func Zero() {
	var (
		i int
		f float64
		b bool
		s string
	)

	fmt.Printf("%v,%v,%v,%q\n", i, f, b, s)
}

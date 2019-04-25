package stringutil

import "fmt"
import "math"

func pow1(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim

}

func IAS() {
	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)
}

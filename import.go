package stringutil

import (
	"fmt"
	"math"
)

func Import() {
	fmt.Printf("Now you have %.1g problems.\n", math.Nextafter(2, 3))
	fmt.Println(math.Pi)
}

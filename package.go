package stringutil

import (
	"fmt"
	"math/rand"
)

func Package() {
	s := rand.Intn(10)
	fmt.Printf("My favorite number is %d\n", s)
}

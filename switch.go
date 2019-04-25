package stringutil

import "fmt"
import "runtime"

func SW() {
	fmt.Print("GO runs on ")
	switch os := runtime.GOOS; os {
	case "drawin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os)
	}
}

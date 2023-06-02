package basic_flow

import (
	"fmt"
)

func Defer() {
	defer fmt.Println("second")

	defer fmt.Println("first")
}

func Stacking_Defers() {
	fmt.Println("counting")

	defer fmt.Println("done")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

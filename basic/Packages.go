package basic

import (
	"fmt"
	"math"
	"math/rand"
)

func Packages() {
	fmt.Println("My favorite number is", rand.Intn(20))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println(math.Pi) //In Go, a name is exported if it begins with a capital letter.
}

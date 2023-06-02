package basic

import "fmt"

const (
	Big = 1 << 100

	Small = Big >> 99
	Pi    = 3.14
)

func needInt(x int) int { return x*10 + 1 }

func needFloat(x float64) float64 {
	return x * 0.1
}

func Constants() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}

func Numeric_Constants() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

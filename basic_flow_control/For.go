package basic_flow

import "fmt"

func For() {
	sum := 0
	for i := 1; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum = ", sum)

	var i = 1
	for i < 10 {
		fmt.Println(i, "**2 = ", i*i)
		i++
	}
}

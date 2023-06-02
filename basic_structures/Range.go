package basic_structers

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func Range() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
	for i, _ := range pow {
		fmt.Printf("%d\t", i)
	}
	fmt.Println()
	for _, v := range pow {
		fmt.Printf("%d\t", v)
	}
	fmt.Println()
}

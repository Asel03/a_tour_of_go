package basic

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var c, python, java bool

func Functions() {
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(25))
	var i int
	fmt.Println(i, c, python, java)
}

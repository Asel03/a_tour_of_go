package basic_structers

import (
	"fmt"
	"strings"
)

func Slices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	var s1 []int = s[1:2]
	var s2 []int = s[2:5]
	var s3 []int = primes[2:5]
	var s8 []int = primes[0:2:5]
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s8)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // Создаем срез "a" из первых двух элементов массива
	b := names[1:3] // Создаем срез "b" из второго и третьего элементов массива
	fmt.Println(a, b)

	b[0] = "XXX" // Изменяем первый элемент среза "b"
	fmt.Println(a, b)
	fmt.Println(names)

	st := []int{2, 3, 5, 7, 11, 13}

	st = st[1:4]
	fmt.Println(st)

	st = st[:2]
	fmt.Println(st)

	st = st[1:]
	fmt.Println(st)

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

package basic_structers

import "fmt"

func Slice_len_cap() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Обрежем срез, чтобы установить его длину равной нулю.
	s = s[:0]
	printSlice(s)

	// Увеличим его длину.
	s = s[:4]
	printSlice(s)

	// Уберем первые два значения.
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

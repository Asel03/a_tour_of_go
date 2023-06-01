package basic_structers

import "fmt"

func Arrays() {
	var a [2]string
	a[0] = "Hello"          // Присваивание значения "Hello" элементу с индексом 0 массива `a`
	a[1] = "World"          // Присваивание значения "World" элементу с индексом 1 массива `a`
	fmt.Println(a[0], a[1]) // Вывод элементов массива `a` с индексами 0 и 1
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13} // Объявление массива `primes` типа `int` с заданными значениями
	fmt.Println(primes)
}

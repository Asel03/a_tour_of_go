package basic_structers

import "fmt"

func Pointers() {
	i, j := 42, 2701

	p := &i // Создание указателя p, который указывает на переменную i
	fmt.Println(*p)
	*p = 21 // Изменение значения переменной, на которую указывает p (i)
	fmt.Println(i)

	p = &j       // Изменение указателя p, чтобы он указывал на переменную j
	*p = *p / 37 // Изменение значения переменной, на которую указывает p (j)
	fmt.Println(j)
}

func exam() {
	a := 100
	b := &a // Создаем указатель "b" и присваиваем ему адрес переменной "a"
	*b++    // Увеличиваем значение, на которое указывает указатель "b", на 1
	c := &b // Создаем указатель "c" и присваиваем ему адрес указателя "b"
	**c++   // Увеличиваем значение, на которое указывает указатель "c", на 1
	*b++    // Увеличиваем значение, на которое указывает указатель "b", на 1
	fmt.Println(a)
}

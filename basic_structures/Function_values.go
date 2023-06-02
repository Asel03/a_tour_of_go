package basic_structers

import (
	"fmt"
	"math"
)

/*
(fn func(float64, float64) float64) - определяет параметр функции fn,
который является функцией, принимающей два аргумента типа float64 и
возвращающей значение типа float64.
Этот параметр fn будет использоваться внутри функции compute.
float64 - указывает, что функция compute будет возвращать значение типа float64.
*/
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func Function_Values() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}

package basic_flow

import (
	"fmt"
	"time"
)

func If() {
	a := 6
	b := 7
	if a < b {
		fmt.Println("a меньше b")
	} else if a > b {
		fmt.Println("a больше b")
	}
}

func Switch() {
	i := 5
	switch i {
	case 0:
		fmt.Println(i, "= Zero")
	case 1:
		fmt.Println(i, "= One")
	case 2:
		fmt.Println(i, "= Two")
	case 3:
		fmt.Println(i, "= Three")
	case 4:
		fmt.Println(i, "= Four")
	case 5:
		fmt.Println(i, "= Five")
	default:
		fmt.Println(i, "= Unknown Number")
	}
}

func Switch_True() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Доброе утро!")
	case t.Hour() < 18:
		fmt.Println("Добрый день!")
	default:
		fmt.Println("Добрый вечер!")
	}
}

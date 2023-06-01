package basic_structers

import "fmt"

// Определение структуры Vertex
type Vertex struct {
	X int
	Y int
}

// Инициализация переменных типа Vertex
var (
	v1 = Vertex{1, 2}  // Создание экземпляра Vertex с явным указанием значений полей
	v2 = Vertex{X: 1}  // Создание экземпляра Vertex с указанием значения только для поля X
	v3 = Vertex{}      // Создание экземпляра Vertex с значениями полей по умолчанию (0 для int)
	z  = &Vertex{1, 2} // Создание указателя на экземпляр Vertex
)

func Struct() {
	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2} // Создание экземпляра Vertex и присваивание его переменной v
	p := &v           // Создание указателя на экземпляр Vertex и присваивание его переменной p
	p.X = 1e9         // Изменение значения поля X через указатель
	fmt.Println(v)
	v.X = 4 // Изменение значения поля X напрямую
	fmt.Println(v.X)
	fmt.Println(v1, z, v2, v3)
}

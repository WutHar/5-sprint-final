package personaldata

import "fmt"

// Personal структура для хранения персональных данных
type Personal struct {
	Name   string
	Weight float64
	Height float64
}

// Print выводит персональные данные
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}

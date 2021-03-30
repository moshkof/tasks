package glava1

import "fmt"

var x int

func Chislo() {
	x = 0
	fmt.Println("Введите число")
	fmt.Scan(&x)
	fmt.Println("Вы ввели число: ", x)
}

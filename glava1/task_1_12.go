package glava1

import "fmt"

var nameTwo string

func NameTwo1() {
	fmt.Println("Введите имя")
	fmt.Scan(&nameTwo)
	fmt.Println("Привет, " + nameTwo + "!")
}

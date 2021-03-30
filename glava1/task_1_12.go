package glava1

import "fmt"

var nameTwo string

func Task_1_12() {
	fmt.Println("Введите имя")
	fmt.Scan(&nameTwo)
	fmt.Println("Привет, " + nameTwo + "!")
}

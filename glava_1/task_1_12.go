package nameTwo1

import "fmt"

var nameTwo string

func nameTwo1() {
	fmt.Println("Введите имя")
	fmt.Scan(&nameTwo)
	fmt.Println("Привет, " + nameTwo + "!")
}

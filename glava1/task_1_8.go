package glava1

import "fmt"

//We display the entered number on the screen
func Task_1_8() {
	var x int = 0
	fmt.Println("Введите число")
	fmt.Scan(&x)
	fmt.Println("Вы ввели число: ", x)
}

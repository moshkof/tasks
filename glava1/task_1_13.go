package glava1

import "fmt"

func Task_1_13() {
	var x int
	fmt.Println("Введите число:")
	fmt.Scan(&x)
	fmt.Println("Следующее за числом " + fmt.Sprint(x) + " число " + fmt.Sprint((x + 1)))
	fmt.Println("Для числа " + fmt.Sprint(x) + "  предыдущее число " + fmt.Sprint((x - 1)))

}

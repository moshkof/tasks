package glava1

import "fmt"

func Summa() {
	var x int
	fmt.Scan(&x)
	fmt.Println("Следующее за числом " + fmt.Sprint(x) + " число " + fmt.Sprint((x + 1)))
	fmt.Println("Для числа " + fmt.Sprint(x) + "  предыдущее число " + fmt.Sprint((x - 1)))

}

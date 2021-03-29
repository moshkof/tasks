package main

import "fmt"

var x int

func main() {
	x = 0
	fmt.Println("Введите число")
	fmt.Scan(&x)
	fmt.Println("Вы ввели число: ", x)
}

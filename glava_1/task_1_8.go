package main

import "fmt"

var x int

func chislo() {
	x = 0
	fmt.Println("Введите число")
	fmt.Scan(&x)
	fmt.Println("Вы ввели число: ", x)
}

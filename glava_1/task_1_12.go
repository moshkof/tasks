package main

import "fmt"

var name string

func main() {
	fmt.Println("Введите имя")
	fmt.Scan(&name)
	fmt.Println("Привет, " + name + "!")
}

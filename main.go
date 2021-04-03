package main

import (
	"fmt"

	"github.com/moshkof/tasks.git/glava1"
	"rsc.io/quote"
)

// Solving problems from the collection of 1400 programming problems
func main() {
	var phrase string = "GO рулит!"
	var name string = "Sergey"
	fmt.Println(phrase)
	fmt.Println(quote.Go())
	// glava1.Task_1_12()
	// glava1.Task_1_13()
	// glava1.Task_1_8()
	// glava1.Task_1_10()
	glava1.Hello(name)
}

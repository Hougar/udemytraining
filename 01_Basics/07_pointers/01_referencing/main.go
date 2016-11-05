package main

import "fmt"

func main() {
	a := 43

	fmt.Println("the variable", a)
	fmt.Println("the memory address", &a)

	var b = &a
	// b is the pointer to the memory address a
	// b is of the typer "int pointer  "

	fmt.Println("the pointer to the memory address a", b)
}

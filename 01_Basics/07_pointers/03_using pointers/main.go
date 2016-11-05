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
	fmt.Println("The value of b what is being stored ", *b)
	// adding a *b to a pointer gives you the value of the pointer not the memory address - dereferncing
	*b = 42
	// the value at address *b change to 42'
	fmt.Println("changed to 42 since the memory adress of a was changed using b", a)
}

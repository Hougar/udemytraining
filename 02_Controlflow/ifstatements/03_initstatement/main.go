package main

import "fmt"

func main() {
	b := true
	if food := "Chocolate"; b {
		// this is the initialization statement
		// b - is the expression of the if statement
		// initialization done bettween food := "Chocolate"
		// inits wihin IF to keep scope within {s}
		// ; seperates the expression from the init

		fmt.Println(food)
	}
}

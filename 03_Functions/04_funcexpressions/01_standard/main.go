package main

import "fmt"

// the ONLY way to have a function inside another function is to assign a function as a variable to an anynomous func

func main() {
	greeting := func() {
		fmt.Println("OLA WORLDO")
	}
	greeting()
	// func greeting is declared here and used above

	fmt.Printf("%T", greeting)
	// func greeting is of type func!
}

package main

import "fmt"

func main() {
	greet("Jane")
	greet("Jack")
}

// greet is called above -> passing in a argument - "Jane" Or "Jack" - they are both names of type string!

func greet(name string) {
	fmt.Println(name)
}

//greet:
// name - identifyer = the variable name i.e. x
// string - type
// name string =  1 parameter

// greet is declared with a parameter name stirng

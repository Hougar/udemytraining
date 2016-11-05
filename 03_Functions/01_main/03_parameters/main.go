package main

import "fmt"

func main() {
	greet("Jack", "Bauer")
}

func greet(fname string, lastname string) {
	fmt.Println(fname, lastname)
}

// Functions put code into a "container or module" and use over and over again !
// variable vs. function - variable no parens - function has parens

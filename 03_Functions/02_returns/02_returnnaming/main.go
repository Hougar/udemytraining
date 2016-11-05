package main

import "fmt"

func main() {
	fmt.Println(greet("jane", "Doe"))
}

func greet(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}

// worse since less redeable i.e.e mutiple parens
// = assigns
// =: declare and initialize

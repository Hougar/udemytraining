package main

import "fmt"

func makegreeter() func() string {
	// func makegreeter()l func () string - anonomys func is of type strung and return it
	return func() string {
		// second half the return what is func string doing?
		return "HELLOS SHWORLD"
	}
}
func main() {
	greet := makegreeter()
	fmt.Println(greet())
}

package main

import "fmt"

var a = 2

func main() {
	if a == 4 {
		fmt.Println("4")
	} else if a == 42 {
		fmt.Println("42")
	} else {
		fmt.Println("Not 42,4")

	}

}

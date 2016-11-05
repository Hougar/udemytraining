package main

import "fmt"

var a = 2

func main() {
	if a == 4 {
		fmt.Println("4")
	} else if a == 42 {
		fmt.Println("42")
	} else if a == 41 {
		fmt.Println("41")
	} else if a == 40 {
		fmt.Println("40")
	} else {
		fmt.Println("Not 42,41,40,4")

	}

}

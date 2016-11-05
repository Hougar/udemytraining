package main

import "fmt"

func main() {
	if !true {
		fmt.Println("NOT TRUE = FALSE")
	}
	if !false {
		fmt.Println("NOT FALSE = TRUE")
		// this runs since true is default
	}
}

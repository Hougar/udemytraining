package main

import "fmt"

func main() {
	letter := rune("Asshole"[0])
	// letter equals a function rune with parameter string A at position 0
	// string is a slice of bytes
	// for string A give me first position
	// rune is an alias for int32
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)
}

// same result as 02_runearenumbers

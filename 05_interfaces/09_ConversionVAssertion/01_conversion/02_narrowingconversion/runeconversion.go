package main

import "fmt"

func main() {
	fmt.Println(string([]byte{'x'}))
	fmt.Println(string([]byte{'y'}))
	// cannot convert rune directly to string but can convert rune - byte slice - to string
	// cannot create a variable = rune and then call --- maybe?

}

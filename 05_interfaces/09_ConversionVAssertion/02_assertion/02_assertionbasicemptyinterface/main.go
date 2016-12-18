package main

import "fmt"

func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string)
	// cane doe this since name is a variable of type empty interface
	// the concrete type which is implemeting the interface is of type string
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string")
	}
}

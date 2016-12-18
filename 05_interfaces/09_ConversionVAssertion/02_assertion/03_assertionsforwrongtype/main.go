package main

import "fmt"

func main() {
	var name interface{} = 2
	str, ok := name.(string)
	// not a string since 2 is int therefore prints else
	// the type we are asserting it to be is on the RIGHT side
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("value is not a string")
	}
}

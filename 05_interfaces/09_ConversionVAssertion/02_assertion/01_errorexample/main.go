package main

import "fmt"

func main() {
	name := "Sydney"
	str, ok := name.(string)
	// non-interface type string on left
	// trying to assertion NOT on an interfaces does not work
	// Assertion = variable.(the type iam asserting it to be)
	if ok {
		fmt.Printf("%q\n", str)
	} else {
		fmt.Printf("value is not a string")
	}
}

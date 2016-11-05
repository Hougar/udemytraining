package main

import "fmt"

func zero(x int) {
	x = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x)
	// x is still 5 with no pointer
	// the value is being passed from function zero i.e. x is being passed not 0
}

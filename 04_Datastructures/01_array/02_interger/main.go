package main

import "fmt"

func main() {
	var x [256]int
	fmt.Println(len(x))
	// = 256
	fmt.Println(x[42])
	// 0 - since default value for int
	for i := 0; i < 256; i++ {
		x[i] = i
		// changes the elements of the array to i based on the loop to 0 - 255 ints
	}

	for _, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}
	// loops over the range of x = 256 prints the index, value, & binary value of items v
	// v = range of x ( 1 iteration of x )
}

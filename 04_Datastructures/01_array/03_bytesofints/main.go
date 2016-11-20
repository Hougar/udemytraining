package main

import "fmt"

func main() {
	var x [256]byte
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {
		x[i] = byte(i)

	}

	for _, v := range x {
		fmt.Printf("%v - %T - %b\n", v, v, v)
	}
	// byte is an alias for uint8 - type
	//uint8 - unassigned type 0 - 255 ( 0 counts 2 ^ 8)
	//int8  the set of all assigned 8 bit intergers -128 - 127 - 256 possible values
}

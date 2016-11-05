package main

import "fmt"

func zero(z int) {
	fmt.Printf("memory of z in func zero = %p\n", &z)
	// the address in func zero of x
	z = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	fmt.Println("memory of x in func main =", &x)
	// address in func main of x
	zero(x)
	fmt.Println("value of x =", x)
	// x is still 5
}

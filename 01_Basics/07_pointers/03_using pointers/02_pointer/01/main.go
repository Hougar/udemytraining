package main

import "fmt"

func zero(z *int) {
	//  z at this point is = 5 but since *int know to take memory address x - since &x
	*z = 0
	// dereferencing zero x i..e changing the meory address of x from X:=5 to =0
}

func main() {
	x := 5
	zero(&x)
	//  pass over by value the memory address of x
	fmt.Println(x)
	//  x is now zero
}

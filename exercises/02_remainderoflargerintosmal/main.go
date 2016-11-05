package main

import "fmt"

var a int
var b int

func main() {
	fmt.Println("Enter small interger: ")
	fmt.Scan(&a)
	fmt.Println("enter larger interger: ")
	fmt.Scan(&b)
	if b%a == 0 {
		fmt.Println("Even - no remainder = 0")
	} else {
		fmt.Println("small number value:", a, "larger number value:", b, "remainder of large/small: ", b%a)
	}
}

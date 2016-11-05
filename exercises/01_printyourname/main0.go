package main

import "fmt"

var a string

func main() {
	fmt.Print("Enter Name: ")
	fmt.Scan(&a)
	fmt.Println("Hello", a)
}

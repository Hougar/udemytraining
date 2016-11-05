package main

import "fmt"

const (W
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

func main() {
	fmt.Println("binary\t\t\tdecimal")
	fmt.Printf("%b\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t ", MB)
	fmt.Printf("%d\n", MB)
}

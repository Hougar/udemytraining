package main

import "fmt"

func main() {
	fmt.Println(half(2))
}

func half(number1 int) (int, bool) {
	return number1 / 2, number1%2 == 0
}

package main

import "fmt"

var x = 0

func wrapper() func() int {

	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(x)
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())

}

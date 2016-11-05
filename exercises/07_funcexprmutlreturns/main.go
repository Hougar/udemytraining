package main

import "fmt"

func main() {

	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}
	fmt.Println(half(1))
	// do this to restrit the scope of half into func main & not at package level
}

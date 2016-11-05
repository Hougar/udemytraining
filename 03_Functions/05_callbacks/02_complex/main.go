package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	// filter has number paramater as type slice of int
	// also hasl callback of type function % bool
	// return a slice of int
	var xs []int
	// var xs is of type slice of int same as xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	fmt.Printf("%T", callback)
	return xs

}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {

		return n > 1
	})
	// xs equals function filter
	// pass in arguments A) slice of intergers i,2,3,4 into filter &
	// pass in b) function with variable n which is an n return a bool as true when n > 1 i.e. 2,3,4,)

	fmt.Println(xs)

}

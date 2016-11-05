// variadic - the final parameter in a function signatures may have the type prefixed with ...
// A function with such a parameter is called variadic and may be invoked with zero or more arguments
// for that parameter

package main

import "fmt"

func main() {
	n := average(4, 3, 2, 1)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	// I can pass in as many float64 arguments as i wnat i.e. th slice of main have 4 arguments - 4,3,2,1
	total := 0.0
	for _, v := range sf {
		total += v
		// += same as total = total + v -> short hand
	}
	return total / float64(len(sf))
}

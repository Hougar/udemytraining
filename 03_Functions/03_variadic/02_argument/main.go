package main

import "fmt"

func main() {
	data := []float64{6, 4}
	n := average(data...)
	// variadic argument - data is of type slice ( not float64) therefore here is first argument 6 then here is second argument 4
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	// variadic parameter
	total := 0.0
	for _, v := range sf {
		total = total + v
	}
	return total / float64(len(sf))
}

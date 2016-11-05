package main

import "fmt"

func main() {
	greatest := max(1, 2, 3, 4)
	fmt.Println(greatest)
}

func max(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		// range loop - loops over a list of numbers
		// gives back two returns the blank identifier _ and v
		// _ is the key - the index  , V is the map the value

		if v > largest {
			largest = v
		}
	}
	return largest
}

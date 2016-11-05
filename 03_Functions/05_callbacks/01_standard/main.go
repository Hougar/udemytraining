// a call back is passing a func as an argument
// ie. you can return a function OR you can call back a function as an argument
package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	// func visit has 2 paramaters
	// paramater 1 - numbers is a slice of int i.e. more than 1 interger
	// paramter 2 = callback is of type function - who has 1 parameter which is int
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
}

package main

import "fmt"

func main() {
	var x [58]string
	// var x is of type array string which stores 58
	fmt.Println(x)
	// [ only
	fmt.Println(len(x))
	// lengh of the array = 58
	fmt.Println(x[42])
	// access an item at position 42 - 43rd element since 0 is first element - nothing assigned so blank
	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}
	fmt.Println(x)
	// prints UTF from 1 - 60 characters
	fmt.Println(len(x))
	// lenghts is unchanged sinc array
	fmt.Println(x[0])
	// prints out A - first letter in UTF-& string(i) for loop
	x[0] = "bananamilkfuckface"
	// changes what is at position 0
	fmt.Println(x[0])
}

// len -gives the length of something - len is built into golang not imported

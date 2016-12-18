package main

import "fmt"

func main() {
	m := make(map[string]int)
	// declare & initializing the map m
	// declare the key - the key is of type string
	// int is the element type of the map i..e the value
	// slices index access with by int i.e. position maps index access by anything
	m["k1"] = 7
	m["k2"] = 13
	// index access the map by using strings k1 & k2
	// the key k1 has a value of 7 the key k2 has a value of 13

	delete(m, "k2")
	// delete key k2 from map m - built in
	fmt.Println("map:", m)
	_, ok := m["k2"]
	// used blank identifier - optional second return which returns the value i.e.deleted so nil - k2 was deleted
	// used the blank id what is the value at k2 - nil since deleted could assign another variable i.e. v to print it out
	fmt.Println("ok?:", ok)
}

package main

import "fmt"

func main() {
	myslice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(myslice)
	fmt.Println(myslice[2:4])  // slicing a slice - 2 <= 4  - c d  (2 3)
	fmt.Println(myslice[2])    // index access  - accessing by index  - c (0 1 2 )
	fmt.Println("mystring"[2]) // index access on a string - 115 UTF position of s
	// a string is a slce of bytes ! - therefore can access the index of UTF8
}

package main

import "fmt"

func main() {
	m := make([]string, 1, 25)
	// m is a slice with length 1 and capacity 25?
	fmt.Println(m)
	// only prints []  - the slice is empty

	changeme(m)
	// slice is a reference type therefore did not need to change memory address as with & int etc - primitiv types
	fmt.Println(m)
}

func changeme(z []string) {
	z[0] = "Craig"
	fmt.Println(z[0])

}

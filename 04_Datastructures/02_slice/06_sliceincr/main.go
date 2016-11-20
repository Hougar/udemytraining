package main

import "fmt"

func main() {
	myslice := make([]int, 1)
	// length and capacity of the underlying array is 1 i.e onlh position 0
	fmt.Println(myslice[0])
	myslice[0] = 7
	// use append if want to add to the capcity i.e. add another int in the slice
	fmt.Println(myslice[0])
	myslice[0]++
	// adds 1 to the interger at position 0
	fmt.Println(myslice[0])
	myslice = append(myslice, 42)
	fmt.Println(myslice[1])
	// append originaly defined myslice as only  have leng/capcity of 1 use append if want to exceed org capcity
	// slices have the ability to grow and shrink
	// myslice[0] = myslice + 1  same as myslice[0]++ same as myslice[0] += 1
	myslice = append(myslice[:1])
	// position 1 is deleted therefore 42 no longer will print
	fmt.Println(myslice[0])

}

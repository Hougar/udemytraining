package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age)
	// prints the memory address of the stored value age

	changeme(&age)
	fmt.Println(&age)
	// memory address of age the same
	fmt.Println(age)
	// the value stroed at the memory address changed with func changeme
}

func changeme(x *int) {
	// change me x is of type pointer interger
	// x receives &age from above
	fmt.Println(x)
	// &age is a memory address so print x is a memory address - refrencing
	fmt.Println(*x)
	// the pointer to x prints what is stored at &age - dereferencmg
	*x = 22
	fmt.Println(x)
	fmt.Println(*x)
	// THE MEMORY ADDRESS NEVER CHANGES !!

}

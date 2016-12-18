package main

import "fmt"

type foo int

func main() {
	var myage foo
	myage = 26
	fmt.Printf("%T %v \n", myage, myage)

	var yourage int
	yourage = 42
	fmt.Printf("%T %v \n", yourage, yourage)
	// this won't work :
	// fmt.Println(myage + yourage)
	// since of different types foo & int

	fmt.Println(int(myage) + yourage)
	// this workds since my age is converted into an int
}

// ONLY AN EXAMPLE BAD PRACTICE TO ALIAS TYPES UNLESS YOU NEED TO ATTACH METHODS - TIME PACKAGE FOR EXAMPLE

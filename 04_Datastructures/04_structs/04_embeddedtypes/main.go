package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

type Doublezero struct {
	Person
	Licensetokill bool
}

func main() {
	p1 := Doublezero{
		// outer type doublezero
		Person: Person{
			first: "James",
			last:  "bond",
			age:   43},
		// inner type person
		Licensetokill: true,
		// Initializing the value of the fields of your struct type
	}
	p2 := Doublezero{
		Person: Person{
			first: "Craig",
			last:  "Beare",
			age:   26,
		},
		Licensetokill: false,
	}
	fmt.Println(p1, p2)
}

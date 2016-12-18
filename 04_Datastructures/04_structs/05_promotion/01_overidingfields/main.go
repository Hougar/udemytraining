package main

import "fmt"

// inner type
type person struct {
	first string // first first
	last  string
	age   int
}

// outer type:
type doublezero struct {
	person
	first         string //  Sseconf first overides the first first - since outer type
	Licensetokill bool
}

func main() {
	p1 := doublezero{

		person: person{
			first: "James",
			last:  "bond",
			age:   43},
		first:         "OVERIDE",
		Licensetokill: true,
	}
	p2 := doublezero{
		person: person{
			first: "Craig",
			last:  "Beare",
			age:   26,
		},
		Licensetokill: false,
	}
	fmt.Println(p1, p2)
	fmt.Println(p1.person.first) // access the overridden first by calling it directly !!

}

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullname() string {
	// (p person) is the receiver
	// the receiver attaches the function fullname to the type person
	return p.first + p.last
}

func main() {
	p1 := person{"Craig", "Beare", 26}
	// the variable p1 is of type person with the values craig, beare , 26
	p2 := person{"Christine", "McPhail", 25}
	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())

}

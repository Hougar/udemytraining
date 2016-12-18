package main

import "fmt"

type person struct {
	first string
	last  string
	age   float64
}

func main() {
	p1 := &person{"Craig", "Beare", 26.6}
	// p1 is of type pointer to person  i.e. *main.person
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.first)
	fmt.Println(p1.age)
}

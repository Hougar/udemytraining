package main

import "fmt"

type animal struct {
	sound string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func species(a interface{}) {
	fmt.Println(a)
}

// func species takes any variable a of type anything ( empty interface)

func main() {
	fido := dog{animal{"woof"}, true}
	meeshka := cat{animal{"meow"}, true}
	species(fido)
	species(meeshka)

}

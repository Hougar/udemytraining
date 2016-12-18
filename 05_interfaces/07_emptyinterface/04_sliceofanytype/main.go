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

func main() {
	fido := dog{animal{"woof"}, true}
	meeshka := cat{animal{"meow"}, true}
	shadow := dog{animal{"woofieee"}, true}
	critters := []interface{}{meeshka, fido, shadow}
	// use empty interface with caution since stricly types natrue of go - ie throwing away benefits of types
	// specific as possible - the bettwe
	fmt.Println(critters)
}

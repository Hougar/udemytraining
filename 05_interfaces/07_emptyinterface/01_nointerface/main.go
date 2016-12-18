// an interface that has NO methods in its definitions
// Eveything implements it since there no signatures !
// go lint = dont secify the type sometimes - i.e var x int = 12 to var x = 12 - more idiomatic
// Exported variables/types capitalized - should have a comment
// linux cmd - ./... go fmt or golint -> do this to everything in directory

// a... -> unlimited amount of variables
// interface{} as a parameter/argument - the method takes an empty inteface - tajes a variable of any type
// every type implements an empty interface  - no methods
// no methods is a subset of the superset methods

// below is NO INTERFACE
package main

import "fmt"

type vehicle struct {
	Seats    int
	Maxspeed int
	Colour   string
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func (v vehicle) Specs() {
	fmt.Printf("Seats %v, Max speed %v. Colour %v\n", v.Seats, v.Maxspeed, v.Colour)
}

func main() {
	prius := car{}
	forte := car{}
	bmwZ37 := car{}
	cars := []car{prius, forte, bmwZ37}

	boeing747 := plane{}
	snowbird := plane{}
	planes := []plane{boeing747, snowbird}

	sanger := boat{}
	laser := boat{}
	boats := []boat{sanger, laser}

	for key, value := range cars {
		fmt.Println(key, "-", value)
	}
	for key, value := range planes {
		fmt.Println(key, "-", value)
	}
	for key, value := range boats {
		fmt.Println(key, "-", value)
	}
}

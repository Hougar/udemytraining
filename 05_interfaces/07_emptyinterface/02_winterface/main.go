// below is WITH INTERFACE

package main

import "fmt"

type vehicles interface {
}

// no methods defined

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

	boeing747 := plane{}
	snowbird := plane{}

	sanger := boat{}
	laser := boat{}

	rides := []vehicles{prius, forte, bmwZ37, boeing747, snowbird, sanger, laser}

	for key, value := range rides {
		fmt.Println(key, "-", value)
	}
	for key, value := range rides {
		fmt.Println(key, "-", value)
	}
	for key, value := range rides {
		fmt.Println(key, "-", value)
	}
	// everything can implement rides methods theres for cant print range rides
}

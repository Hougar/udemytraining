package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c *circle) area() float64 {
	// c circle is a pointer receiver
	// must be an address because go is untyped language - not decided untill used cannot pass a value into a pointer when type has not yet been determined
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("area", s.area())
}

func main() {
	c := circle{5}
	info(&c)
	// this is a pointer VALUE not a pointer type - can go to the receiver since of pointer type
}

package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

// when declaring the shape interface must have the method area()  with signature float64

func (s Square) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius
}

func info(x Shape) {
	fmt.Println(x)
	fmt.Println(x.area())
	// x.area() a method call on the interface value Shape
	// circle & Square run dfferent area methods
}

func main() {
	s := Square{5}
	c := Circle{1}

	info(s)
	info(c)
	// polymorphism - s & c are of different types (circle & Square) but both are Shapes
	// func info has parameters of type shape - which is of underlying type Shapes
	// Concrete type = Circle Or Square
	// the concrete type implementing the shape interface is Square OR Circle
	// knows which method to call since the arguments s,c and parameters in the corespondin methods
	// Can actually have 1 function name be named more than once sort of
}

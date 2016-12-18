package main

import "fmt"

type Square struct {
	side float64
}

// create a user defined type: Square
// Square is of underlying type struct
// Struct has the field side which is of type float64

func (z Square) area() float64 {
	return z.side * z.side
	// the receiver (z Square) is attaching the method of func area() to struct Square
}

type Shape interface {
	area() float64
}

// underlying type of Shape is interface
// an interface defines functionality
// anything with area() float64 implements the interface Shape
// Square implements the shape interface since its signature matches Shape

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

// info takes any shapes - square is a shape therefore type square can be passed through info

func main() {
	s := Square{10}
	info(s)
	fmt.Println()
	fmt.Println(s.area())
	// variable s sine it is of type Square has the method area attached s
}

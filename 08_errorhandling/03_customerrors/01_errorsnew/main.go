package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("norgatem math: imaginary root not appropiate for this code")
		// New is the only function in package errors - takes a text string returns "error"
		// error - is a type it is an interface - Error() string - method - will implement the error interface and be of type error - can return erros.New since returns and error
		// use errors.New to provide very speific information about your errors !
	}
	return 42, nil
}

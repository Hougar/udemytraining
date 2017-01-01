package main

import (
	"errors"
	"fmt"
	"log"
)

var Errnegativesqrt = errors.New("norgatem math: imaginary root not appropiate for this code")

// Idiomatic to have your variable name start with Err
// reuse your error message used in package io & bufio

func main() {
	fmt.Printf("%T\n", Errnegativesqrt)
	// type : *errors.errorString
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, Errnegativesqrt
	}
	return 42, nil
}

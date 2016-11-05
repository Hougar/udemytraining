package main

import "Fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world()
	hello()
}

// func hello prints now before func world
//defer will defer the function until the function it is in exists i.e. main
// open a file want to close it when done ... use defer close since last step is to close !

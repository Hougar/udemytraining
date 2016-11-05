package main

import "Fmt"

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world")
}

func main() {
	world()
	hello()
}

// prints world hello not hello wordl because top to bottom main

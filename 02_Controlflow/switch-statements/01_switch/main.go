package main

import "fmt"

func main() {
	switch "tyler" {
	case "daniel": // COLON REQUIRED HERE
		fmt.Println("FUCK YOU DANIEL")
	case "tyler":
		fmt.Println("FUCK YOU TYLER")
	default:
		fmt.Println("who the FUCK are you?")

	}
}

package main

import "fmt"

func main() {
	switch "tyler" {
	case "daniel", "tyler":
		fmt.Println("FUCK YOU DANIEL OR TYLER")
	case "susan": // Cannot have tyler in OR and as a distinct case in switches - no duplicates
		fmt.Println("FUCK YOU SUSAN")
	default:
		fmt.Println("who the FUCK are you?") // default will print when you have a mutiple eval i.e in case ln7

	}
}

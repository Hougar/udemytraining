package main

import "fmt"

func main() {
	switch "tyler" {
	case "daniel":
		fmt.Println("FUCK YOU DANIEL")
		fallthrough // this one does proceed since before tyler the swtich
	case "tyler":
		fmt.Println("FUCK YOU TYLER")
		fallthrough // default & tyler Println will continue
	default:
		fmt.Println("who the FUCK are you?")

	}
}

package main

import "fmt"

func main() {
	if true && false {
		fmt.Println("TRUE AND FALSE NOT POSSIbLE")
	} else {
		fmt.Println("got to be this one! ")
		// can't be both true and false default therefore else runs
	}
}

package main

import "fmt"

func main() {
	if true && false || false && true || !false && false {
		fmt.Println("RUN")
	} else {
		fmt.Println("DIDNT RUN RUN")
	}
}

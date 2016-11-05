package main

import "fmt"

func main() {
	func() {
		// anonynous no func name
		fmt.Println("IM DRIVING NOW!!!")
	}()
	// self executing - () after excecutes function
}

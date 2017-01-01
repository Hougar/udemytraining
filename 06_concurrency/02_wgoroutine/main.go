package main

import "fmt"

func main() {
	go foo()
	go bar()
}

// nothing prints out when adding go
// Have concurrency
// Three go routines running
// Func main
// foo -spools up
// bar  - spools up
// func main exists nothing pritned

func foo() {
	for i := 0; i < 30; i++ {
		fmt.Println("Foo:", i)
	}
}

func bar() {
	for i := 0; i < 30; i++ {
		fmt.Println("Bar:", i)
	}
}

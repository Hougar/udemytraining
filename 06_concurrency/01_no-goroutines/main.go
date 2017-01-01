package main

import "fmt"

func main() {
	foo()
	bar()
}

// No concurrency
// items run sequentially foo is printed then bar since foo is line 6 and bar line 7
// func main()  calls foo then bar
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

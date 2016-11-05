package main

import "fmt"

func main() {
	a := 0
	for i := 0; i < 10; i++ {
		if i%3 == 0 {
			a += i
		} else if i%5 == 0 {
			a += i
		}
	}
	fmt.Println(a)
}

// += add the variable i when condition is true to counter variable a

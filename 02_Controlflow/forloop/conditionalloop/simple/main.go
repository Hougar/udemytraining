package main

import "fmt"

func main() {
	i := 0
	for i < 10 { // i <10 is the condition - i.e. init; condition; then (post)
		fmt.Println(i)
		i++
	}
}

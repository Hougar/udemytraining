package main

import "fmt"

func main() {
	i := 1
	for {
		i++
		if i%2 == 0 {
			continue // continue if even number
		}
		fmt.Println(i)
		if i >= 50 { // stop if value is greater than OR equal to 50
			break
		}
	}
}

package main

import "fmt"

func main() {
	myMap := [3]int{8282, 8283, 8384}
	for k, v := range myMap {
		fmt.Println(k, v)
		// k prints 0, 1 ,2 since the place in myMap int is 0,1,2
	}
}

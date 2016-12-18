package main

import "fmt"

func main() {
	mygreeting := map[int]string{
		0: "Goodmorning",
		1: "Bon matin",
		2: "Gooten  Morgen",
		3: "Bueonos Matinus",
	}
	fmt.Println(mygreeting)
	// delete(mygreeting, 2)

	if val, exists := mygreeting[2]; exists {
		fmt.Println("That value does exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value does not exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
	fmt.Println(mygreeting)
}

// purpose of comma ok idiom - is the value missing or is it there but it is at zero?

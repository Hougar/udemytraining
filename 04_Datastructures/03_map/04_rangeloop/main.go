package main

import "fmt"

func main() {
	mygreeting := map[int]string{
		0: "Good mornting",
		1: "Bon matin",
		2: "Gooten Morgen",
		3: "Buenos dias",
	}
	for key, v := range mygreeting {
		fmt.Println(key, " - ", v)
	}

}

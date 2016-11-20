package main

import "fmt"

func main() {
	var records [][]string
	// a string of type string start at 0
	// a slice that stores slices of strings
	student1 := make([]string, 4)
	// first student a string of length 4 (4 things stored)
	student1[0] = "Beare"
	student1[1] = "Craig"
	student1[2] = "100.00"
	student1[3] = "92.21"
	records = append(records, student1)
	student2 := make([]string, 4)
	student2[0] = "McPhail"
	student2[1] = "Christine"
	student2[2] = "100.00"
	student2[3] = "94.21"
	records = append(records, student2)
	fmt.Println(records)
}

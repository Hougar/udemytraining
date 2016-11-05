package main

import "fmt"

const meterstoyards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter Meters Swam: ")
	fmt.Scan(&meters)
	yards := meters * meterstoyards
	fmt.Println(meters, " meters is ", yards, "yards")

}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = "12"
	var y = 6
	z, _ := strconv.Atoi(x)
	// strconv.Atol is mutiply value error - for error so need blank identifier to throw away
	fmt.Println(z + y)
}

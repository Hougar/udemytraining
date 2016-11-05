package main

import "fmt"

var t = 7
var v = "STRINGBEANS"

func switchontype(x interface{}) {
	switch x.(type) { // this is not conversion but "assertion" - swich x based on its type "." is the assertion ID?
	case int:
		fmt.Println(t, " - this is an int - interger")
	case string:
		fmt.Println(v, "- this is a string - a whole bunch of runes - a whole bunch of int32")
	default:
		fmt.Println("- if error than print - unknown")
	}
}

func main() {

	switchontype(t)
	switchontype(v)

}

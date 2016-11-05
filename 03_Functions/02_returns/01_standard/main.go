package main

import "Fmt"

func main() {
	s := greet("jane", "doe")
	fmt.Println(s)
}

// greet - by setting equal to variable s can now print ( substituition)

func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}

// return - return - brings fmt.Sprint(fname,lname) - up past the parens to "string"
// Sprint - print string - similiar to concentenate
// ... unlimited amount of argument
// interface{} -> can be of any type

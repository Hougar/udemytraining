package main

import "fmt"

func main() {
	fmt.Println(greet("jane", "DOE"))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)

}

// more than one return since did return as string, string - and gave return - fmt.Sprint 2 Sprints

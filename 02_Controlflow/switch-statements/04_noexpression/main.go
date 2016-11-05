package main

import "fmt"

func main() {
	myfriend := "Daniel"
	switch {
	case len(myfriend) == 5:
		fmt.Println("FUCK YOU TYLER MR. 5 LETTERS IN YOUR NAME")
	case len(myfriend) == 6:
		fmt.Println("FUCK YOU DANIEL MR. 6 LETTERS IN YOUR NAME")
	}
}

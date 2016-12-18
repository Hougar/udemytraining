package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	First        string
	Last         string
	Age          int
	notExpported int
}

func main() {
	var p1 Person
	// use var if you want to initialize to zero value - idiomatic way
	rdr := strings.NewReader(`{"First":"James","Last":"Bond","Age":20}`)
	// rdr is used to similate json coming infrom somewhere=
	json.NewDecoder(rdr).Decode(&p1)
	// New decoder takes a reader - not a string
	// get a pointer to a reader
	// If we have a pointer to reader have the following methods - functions which accept *Reader
	// - use the a method which returns Read []bytes error - implement the read interface
	// New decoder wants a reader!
	// New decoder gives us a pointer to a decoder
	// Have the methods available which accept a *Decoder
	// One of these methods is Decode
	// Decode - the reader & put it into the adress of p1 - our struct
	// "pointed to" is the address
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

}

//encoding - using the writer interface
package main

import (
	"encoding/json"
	"os"
)

type person struct {
	First       string
	Last        string
	Age         int
	Notexported int
}

func main() {
	p1 := person{"James", "Bond", 42, 38383}
	json.NewEncoder(os.Stdout).Encode(p1)
	// enconding and writing to standardout - monitor
	//return a pointer to an encoder ;
	// gain acess to the method .Encode(p1)
	// NewEncoder(os.Stdout) returns a pointer to an encoder
	// polymorphism - one function can take in a whole bunch of different things & implement them in different ways

	// call in NewEncoder
	// givest pointer to the Encoder
	// allows us to call Encode
	// Encode takes v interface & returns an erro if there is one
}

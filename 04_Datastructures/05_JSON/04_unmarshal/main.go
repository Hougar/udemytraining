package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int `json:"wisdomscore"`
	// NEED to be exported - JSON not sure why...

}

func main() {
	var p1 person
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)

	bs := []byte(`{"First":"James", "Last":"Bond", "wisdomscore":42}`)
	// back ticks needed - raw string literal to go into JSON "" the character not the use
	// using string() int() - is conversion so is []byte()
	// wisdom score coming in as a field - converted with tag in struct field descriptions

	json.Unmarshal(bs, &p1)
	// nothing printes since you unmarshaled it!
	// passing in the address of variable p1 to your struct bs

	fmt.Println("----------------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}

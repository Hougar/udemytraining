package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notexported int
} // fields need to be exported

func main() {
	p1 := person{"Craig", "Beare", 26, 42}
	bs, _ := json.Marshal(p1)
	// taking a struct p1 and turning it into a byte slice
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
	// noted that not exported variable does not go through since no CAP
}

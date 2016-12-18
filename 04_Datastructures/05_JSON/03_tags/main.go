package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string `json:"-"`           // does not marshal to JSON
	Age         int    `json:"wisdomscore"` // changes wge to wisdom score for JSON purposes
	notexported int
}

func main() {
	p1 := person{"Craig", "Beare", 26, 42}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
}

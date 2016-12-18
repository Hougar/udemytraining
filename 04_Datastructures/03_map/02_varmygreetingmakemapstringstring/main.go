package main

import "fmt"

func main() {
	var mygreetting = make(map[string]string)
	// key of string element of string
	// var mygreetting make(map[string]string)
	// without the equation sign will make a nil map
	// cant append a nil map like you can a slice

	// var mygrettings = map[string]string{} --- same as above using make! woot (except not nil)
	// {} the curlies make it a composite literal even if nothing was added to it and with the {} the map is no longer nil

	//mygreetting := make(map[string]string) same as above

	mygreetting["Claud"] = "Goten Morgen"
	mygreetting["Mathieu"] = "Bon matin"

	//mygreetting := make(map[string]string) { - same as above composite literal
	// "Claud": "Goten Morgen",
	// "Mathieu": "Bon Matin",
	// }

	// adding in an entry in a composite literal
	// mygreetting["Yoshi"] = "Kaneecheewa"

	fmt.Println(mygreetting)
	// prints entire map
	fmt.Println(mygreetting == nil)
	// map is not nil therefore prints false

}

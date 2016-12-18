package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

// Struct sequenced of named elements called a field  - each of which has a name and typ
// = Struct: {
// 	first string
// 	last  string
// 	age   int
// }
// first string - field
// first - name
// string type
// type can also have tages i.e. ` json - how to put info interacting with - i. e. change the struct when snt out but internally keep the same
// the underlying type is a struct w 3 underlying types
`
// we have created are own user defined type
// the person type has an underlying type of string,string,int
type foo int



func main() {
	p1 := person{"Craig", "Beare", 26}
	p2 := person{"Christine", "McPhail", 25}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	// value of type person assigning to p1
	// {struct literal}
	// p1 is of type person

	var myage foo
	myage = 26
	fmt.Printf("%T %v \n", myage, myage)
	// creating your own type named my age which is an int
	// %T - in package main of type foo - main.foo

}

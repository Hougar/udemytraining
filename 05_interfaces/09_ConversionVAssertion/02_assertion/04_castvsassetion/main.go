package main

import "fmt"

func main() {
	var name interface{} = "22"
	str, ok := name.(string)
	// assertion on right side
	if ok {
		fmt.Printf("%T - Assert\n", str)
	} else {
		fmt.Printf("value is not a string")
	}
	cast(7.24)

}

func cast(x float64) {
	fmt.Printf("%T - no conversion \n", x)
	fmt.Printf("%T - cast conversion\n", int(x))

}

// CANNOT use CONVERSIOn on a interface type i.e. cannot CAST() it must ASSERT.IT

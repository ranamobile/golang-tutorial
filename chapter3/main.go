package main

import "fmt"

func main() {
	// First variable example
	var x string
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	// Second variable example
	x = "first"
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)
}

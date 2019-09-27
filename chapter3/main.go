package main

import "fmt"

func main() {
	var x string
	var y string

	// First variable example
	x = "first"
	fmt.Println(x)
	x = "second"
	fmt.Println(x)

	// Second variable example
	x = "first"
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	// Third variable example
	x = "hello"
	y = "world"
	fmt.Println(x == y)

	y = "hello"
	fmt.Println(x == y)

	// Duck typing example
	z := 5
	fmt.Println(z)
}

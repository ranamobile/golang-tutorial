package main

import "fmt"

var x string = "Hello, World"

func main() {
	const y string = "Goodbye"
	y = "Nope!"  // compiler error

	fmt.Println(x)
	fmt.Println(y)
}

func f() {
	fmt.Println(x)
	fmt.Println(y)  // compiler error
}

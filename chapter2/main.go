package main

import "fmt"

func main() {
	// Number examples
	fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1.0 + 1.0 =", 1.0 + 1.0)

	// String examples
	fmt.Println(len("Hello, World"))
	fmt.Println("Hello, World"[1])
	fmt.Println("Hello, " + "World")

	// Boolean examples
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
}

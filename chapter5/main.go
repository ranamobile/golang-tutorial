package main

import "fmt"

func main() {
	// arrays
	var x [5]int
	x[4] = 100
	fmt.Println(x)

	// slices
	y := make([]float64, 5, 10)
	fmt.Println(y)
	z := []float64{1, 2, 3, 4, 5}
	fmt.Println(z)
	w := append(z, 6, 7)
	fmt.Println(w)

	// maps
	a := make(map[string]int)
	a["one"] = 1
	a["two"] = 2
	fmt.Println(a)

	name, ok := a["one"]
	fmt.Println(name, ok)
}

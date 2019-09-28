package main

import "fmt"

func average(xs []float64) float64 {
	total := add(xs...)
	return total / float64(len(xs))
}

func add(args ...float64) float64 {
	total := 0.0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	testxs := []float64{98, 93, 77, 82, 83}
	fmt.Println(average(testxs))
}

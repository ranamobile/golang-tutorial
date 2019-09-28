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

func first() {
	fmt.Println("1st")
}

func second() {
	fmt.Println("2nd")
}

func main() {
	testxs := []float64{98, 93, 77, 82, 83}

	// basic function call
	avg := average(testxs)
	fmt.Println(avg)

	// nested function declaration
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())

	// defer - called when function completes
	// usually useful for closing a file
	defer second()
	first()

	// panic and recover
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}

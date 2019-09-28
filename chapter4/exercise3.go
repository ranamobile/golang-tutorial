package main

import "fmt"

// Write a program that prints the numbers
// from 1 to 100, but for multiples of three,
// print "Fizz" instead of the number, and
// for multiples of five, print "Buzz." For
// numbers that are multiples of both three
// and five, print "FizzBuzz."
func main() {
	for i := 1; i <= 100; i++ {
		if i % 3 == 0 {
			fmt.Print("Fizz")
		}
		if i % 5 == 0 {
			fmt.Print("Buzz")
		}

		if i % 3 == 0 || i % 5 == 0 {
			fmt.Println()
		} else {
			fmt.Println(i)
		}
	}
}

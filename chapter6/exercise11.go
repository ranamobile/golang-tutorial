package main

import "fmt"

func swap(x *int, y *int) {
	w := *x  // save value of x
	*x = *y  // set value of x to y
	*y = w  // set value of y to original x
}

func main() {
	xPtr := 5
	yPtr := 10

	fmt.Println(xPtr, yPtr)
	swap(&xPtr, &yPtr)
	fmt.Println(xPtr, yPtr)
}

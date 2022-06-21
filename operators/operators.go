package main

import "fmt"

func main () {
	a, b := 2, 4

	// Comparison operator - falsy
	var falsy = a > 3

	fmt.Println("Falsy", falsy)

	// Comparison operator - truthy
	var truthy = (b > 3)
	fmt.Println("Truthy", truthy)
}
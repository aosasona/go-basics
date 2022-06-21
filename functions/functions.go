package main

import "fmt"

// Sum
func sum(a, b int) int {
	return a + b
}

// Average
func average(a, b, c int) int {
	return (a + b + c) / 3
}

func main(){
	var (
		a int = 10
		b int = 20
		c int = 30
	)
	sumValue := sum(a, b)

	fmt.Println("The sum is", sumValue)
	fmt.Println("The average is", average(a, b, c))
	greet("Ayodeji")
}

// Writing functions below the main function
func greet(name string) {
	fmt.Println("Hello", name)
}
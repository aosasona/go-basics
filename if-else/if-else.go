package main

import "fmt"

func main () {
	var age int = 18
	fmt.Println("John is", isAdult(age))
}

// Check if a number is greater than 18
func isAdult(age int) string {
	// return age > 18
	if age > 18 {
		return "an adult"
	} else if age == 18 {
		return "sort of an adult, don't let him drink yet though"
	} else { return "a child" }
}

// Note to self : GO is quite similar to python, it seems to care about indentation a lot when it comes to if-else statements
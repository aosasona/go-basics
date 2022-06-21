package main

import "fmt"

func main(){

	// Using VAR
	var name string = "Ayodeji"
	fmt.Println("Hello, I am",name)

	// Using the create and assign operator
	age := 18
	fmt.Println("I am ",age," years old")

	// Creating multiple variables
	var1, var2 := 2, 5

	// Instantiating a variable
	var sum int
	sum = var1 + var2
	fmt.Println("The sum is", sum)

	// Batch assignment
	var (
		firstName string = "John"
		lastName string = "Doe"
	)
	fmt.Println("This guy is", firstName, lastName)

	// Reassigning a variable
	sum = var2 + 20
	fmt.Println(firstName,"is", sum)
}
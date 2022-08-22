package main

import "fmt"

func main() {

	// doing it the boring way
	var Ages [5]int
	Ages[0] = 10
	Ages[1] = 20
	Ages[2] = 30
	Ages[3] = 40
	Ages[4] = 50

	for i := 0; i < len(Ages); i++ {
		fmt.Println("Age - ", Ages[i])
	}

	// using the var keyword
	var Heights = [...]int{150, 143, 167, 148}
	for i := 0; i < len(Heights); i++ {
		fmt.Println("Height - ", Heights[i])
	}

	// using struct as an array type and direct assignment
	type Person struct {
		name string
		age  int
	}

	people := [...]Person{
		{"Tom", 50},
		{"Jerry", 50},
		{"Alan", 40},
	}

	for i := 0; i < len(people); i++ {
		fmt.Println("Name - ", people[i].name)
		fmt.Println("Age - ", people[i].age)
	}
}

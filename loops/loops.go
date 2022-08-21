package main

import "fmt"

func main() {
	fizzBuzz(50)
}

func fizzBuzz(limit int) {

	i := 0

	for i < limit {
		modThree := i%3 == 0
		modFive := i%5 == 0

		if modThree && modFive {
			fmt.Println("FizzBuzz - ", i)
		} else if modThree {
			fmt.Println("Fizz - ", i)
		} else if modFive {
			fmt.Println("Buzz - ", i)
		}
		i++
	}

}

package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	address   *Address
}

type Address struct {
	city  string
	state string
}

func main() {

	var (
		tom Person = Person{
			firstName: "Tom",
			lastName:  "Cruise",
			age:       50,
			address: &Address{
				city:  "Los Angeles",
				state: "CA",
			},
		}
		jerry Person = Person{
			firstName: "Jerry",
			lastName:  "Seinfeld",
			age:       50,
			address: &Address{
				city:  "New York",
				state: "NY",
			},
		}
	)

	tomString := "Tom is " + strconv.Itoa(tom.age) + " years old and lives in " + tom.address.city + ", " + tom.address.state
	jerryString := "Jerry is " + strconv.Itoa(jerry.age) + " years old and lives in " + jerry.address.city + ", " + jerry.address.state

	fmt.Println(tomString)
	fmt.Println(jerryString)

}

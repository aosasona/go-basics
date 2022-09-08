package main

import (
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {
	people := [...]Person{
		{"Tom", "Cruise"},
		{"Jerry", "Seinfeld"},
		{"Michael", "Landon"},
		{"Alan", "Smith"},
		{"Tony", "Williams"},
		{"John", "Jones"},
	}

	replaceLastName(&people)
	for _, person := range people {
		println(person.firstName, person.lastName)
	}
}

func replaceLastName(people *[6]Person) {
	matches := "smith jones williams"
	i := 0
	for i < len(*people) {
		if strings.Contains(matches, strings.ToLower(people[i].lastName)) {
			people[i].lastName = "Doe"
		}
		i++
	}
}

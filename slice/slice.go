package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	students := []Person{
		{"Tom", "Cruise"},
		{"Jerry", "Seinfeld"},
		{"Michael", "Scoffield"},
	}

	fmt.Println("...Looping...")
	for i := 0; i < len(students); i++ {
		fmt.Println("Name - ", students[i].firstName)
	}

	fmt.Println("...Selecting one student...")
	luckyStudent := students[2]
	fmt.Println("Name - ", luckyStudent.firstName)

	fmt.Println("...Adding more students...")
	students = append(students, Person{"Alan", "Turing"}, Person{"Tony", "Stark"})
	newStudents := students[3:]
	for i := 0; i < len(newStudents); i++ {
		fmt.Println("Name - ", newStudents[i].firstName)
	}

	// merging two slices
	fmt.Println("...Merging two slices...")
	extraSlice := []Person{
		{"John", "Doe"},
		{"Jane", "Doe"},
	}
	combinedStudents := append(students, extraSlice...)
	i := 0
	for i < len(combinedStudents) {
		fmt.Println("Name - ", combinedStudents[i].firstName)
		i++
	}
}

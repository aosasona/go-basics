package main

import "fmt"

func main() {
	result := checkAge(20)
	fmt.Println(result)
}

func checkAge(value int) string {
	switch age := value; {
	case age < 18:
		return "You are a child"
	case age >= 18:
		return "You are eligible to drink"
	case age <= 60:
		return "You are eligible to retire, stop drinking"
	default:
		return "Sorry, IDK"
	}
}

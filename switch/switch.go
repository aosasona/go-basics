package main

import "fmt"

func main() {
	result := checkAge(20)
	fmt.Println(result)
}

func checkAge(age int) string {
	switch age {
	case 18:
		return "You are eligible to drink"
	case 21:
		return "You haven't had your first drink?"
	case 60:
		return "You are eligible to retire, stop drinking"
	default:
		return "Sorry, IDK"
	}
}

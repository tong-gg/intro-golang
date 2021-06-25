package main

import "fmt"

func main() {
	var day string = "Monday"
	fmt.Println("day: ", day)

	day = "Tuesday"
	fmt.Println("day: ", day)

	const dayOfBirth = "Wednesday"
	fmt.Println("Tong's day of birth: ", dayOfBirth)
	// dayOfBirth = "Thursday" <--- error right here! cannot re-assign const variable

	const (
		Sunday = iota // auto increment number from 0
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday: ", Sunday)
	fmt.Println("Monday: ", Monday)
	fmt.Println("Tuesday: ", Tuesday)
	fmt.Println("Wednesday: ", Wednesday)
	fmt.Println("Thursday: ", Thursday)
	fmt.Println("Friday: ", Friday)
	fmt.Println("Saturday: ", Saturday)
}

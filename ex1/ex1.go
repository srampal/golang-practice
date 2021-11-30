// Ex1 is the first exercise package
package ex1

import "fmt"

// Exercise runs this exercise
func Exercise() {
	fmt.Printf("\nRunning Exercise 1\n\n")

	// Basic type definitions

	type age uint8
	type dayOfWeek string

	const (
		Monday    dayOfWeek = "Monday"
		Tuesday             = "Tuesday"
		Wednesday           = "Wednesday"
		Thursday            = "Thursday"
		Friday              = "Friday"
		Saturday            = "Saturday"
		Sunday              = "Sunday"
	)

	type person struct {
		name   string
		years  age
		dayOff dayOfWeek
	}

	var teamA = []person{{"John", 45, Sunday}, {"Jack", 53, Friday}, {"Jill", 47, Saturday}}

	teamB := []person{{"Joe", 48, Tuesday}, {"Mary", 44, Thursday}}

	fmt.Printf("Type of teamA is %T value is %v  or equivalently %#v\n\n", teamA, teamA, teamA)

	fmt.Print("teamB has ... \n")
	for i, val := range teamB {
		fmt.Printf("%d) %v\n", i, val)
	}
} // main()

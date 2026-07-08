package main

import (
	"fmt"
	"time"
)

func main() {
	// switch statement
	// you dont need to break after each case
	// the default case is optional
	switch 1 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("other")
	}

	// switch statement with multiple cases
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("Sunday/Saturday")
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Monday-Friday")
	default:
		fmt.Println("other")
	}

	// switch statement with type switch
	typeChecker := func(i interface{}) string {
		switch i.(type) {
		case int:
			return "int"
		case string:
			return "string"
		default:
			return "other"
		}
	}
	fmt.Println(typeChecker(1))
	fmt.Println(typeChecker("hello"))
	fmt.Println(typeChecker(true))
}

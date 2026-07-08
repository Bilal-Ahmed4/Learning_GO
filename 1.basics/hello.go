package main

import "fmt"

func main() {

	// if you declare a variable than you have to initliate it and use it
	// you have to assign the type if you are not initializing it with a value
	var name string
	name = "Bilal"

	fmt.Println("Hello, World!", name)

	if 1 == 1 {
		fmt.Println("yes")

	} else {
		fmt.Println("no")
	}

	// we can declare variables inisde the if statement and can use them inside if and else blocks
	// if we declare or intiliaze inside the else if like this than it can be used inside the that else if block only
	// else cannt use the variable declared inside the if block
	if i := 10; i > 18 {
		fmt.Println("yes", i)

	} else {
		fmt.Println("no", i)
	}
}

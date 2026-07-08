package main

import "fmt"

// := cannot be used here — only var works at package scope.

func main() {
	//var with explicit type
	var name string
	var age int = 25
	fmt.Println(name, age)

	//var with type inference (no explicit type)
	var name2 = "Alice"
	var age2 = 25
	fmt.Println(name2, age2)

	name3 := "Alice"
	age3 := 25
	//Only works inside functions, not at package level.
	fmt.Println(name3, age3)

	//Declaring multiple variables at once
	var a, b, c int
	var x, y = 10, "hello"
	w, q := 1, 2
	fmt.Println(a, b, c, x, y, w, q)

	// Grouped var block
	var (
		name4   string = "Alice"
		age4    int    = 25
		isAdmin bool   = true
	)
	fmt.Println(name4, age4, isAdmin)

	//Zero value declaration (no initializer)
	var count int    // 0
	var name5 string // ""
	var active bool  // false
	var list []int   // nil
	fmt.Println(count, name5, active, list)

	//:= cannot be used here — only var works at package scope.

}

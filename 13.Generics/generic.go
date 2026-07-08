package main

import (
	"fmt"
)

// generics are used when we have to pass the mulitple type
// we can also T interfac{} equivalent to T any
// also can use T int | string |..
// also can use T comaparable(provide common types )
func printSlice[T comparable](element []T) {
	for _, item := range element {
		fmt.Println(item)
	}
}

// we can also use the generic to make the struct

type order[T any] struct {
	element []T
}

func main() {
	var element []int = []int{1, 3, 3}
	lang := []string{"golang", "c++", "java"}

	printSlice(element)
	printSlice(lang)

}

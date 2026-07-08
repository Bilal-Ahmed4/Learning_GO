package main

import "fmt"

// variadic function can take a varibale or any number of arguments println is example of variadic function
func sum(nums ...int) int {
	// it will give as a slice and we can loop through it
	total := 0
	for _, num := range nums {
		fmt.Print(num, " ")
		total += num
	}
	return total
}

func main() {
	// result := sum(1, 2, 30, 4, 5)
	// can also pass it through a slice
	nums := []int{1, 2, 30, 4, 5}
	num2 := []int{}
	num2 = append(num2, 5)
	result := sum(nums...)
	fmt.Println(result)
	fmt.Println(num2)
}

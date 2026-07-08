package main

import "fmt"

func main() {
	// traditional for loop
	println("Traditional for loop")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// while loop equivalent
	fmt.Println("While loop equivalent")
	j := 1
	for j < 10 {
		// you can also use the break and the continue
		fmt.Println(j)
		j++
	}

	// infinti loop
	// for {
	// 	 print anything
	// }
	fmt.Println("for range ")
	for i := range 10 {
		fmt.Println(i)
	}
}

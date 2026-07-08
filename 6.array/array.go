package main

import "fmt"

func main() {
	var arr [4]int

	fmt.Print(len(arr))
	fmt.Println(arr)

	// also
	num := [4]int{1, 2, 3, 4}
	fmt.Println(num)

	// 2d array 2x2
	num1 := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println(num1)
}

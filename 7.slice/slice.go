package main

import "fmt"

func main() {

	// ============================================================
	// 1. WHAT IS A SLICE?
	// A slice is a flexible, resizable "view" into an array.
	// Unlike arrays, slices don't have a fixed size.
	// ============================================================

	fruits := []string{"apple", "banana", "cherry"}
	fmt.Println("1. Initial slice:", fruits)
	fmt.Println("   Length:", len(fruits), "Capacity:", cap(fruits))

	// ============================================================
	// 2. SLICE FROM AN ARRAY
	// array[low:high] -> low is inclusive, high is exclusive
	// ============================================================

	numbersArray := [6]int{10, 20, 30, 40, 50, 60}
	sliceFromArray := numbersArray[1:4] // indices 1,2,3
	fmt.Println("\n2. Array:", numbersArray)
	fmt.Println("   Slice [1:4]:", sliceFromArray)

	// Slices share memory with their source array!
	sliceFromArray[0] = 999
	fmt.Println("   Original array changed too:", numbersArray)

	// ============================================================
	// 3. LENGTH vs CAPACITY
	// len() = elements currently in the slice
	// cap() = elements from slice start to end of underlying array
	// ============================================================

	nums := []int{1, 2, 3, 4, 5}
	sub := nums[1:3]
	fmt.Println("\n3. Sub-slice [1:3]:", sub)
	fmt.Println("   len(sub):", len(sub), "cap(sub):", cap(sub))

	// ============================================================
	// 4. CREATING SLICES WITH make()
	// make([]Type, length, capacity)
	// ============================================================

	madeSlice := make([]int, 3, 5)
	fmt.Println("\n4. make([]int, 3, 5):", madeSlice)

	// ============================================================
	// 5. APPENDING TO A SLICE
	// append() returns a (possibly new) slice header
	// ============================================================

	colors := []string{"red", "green"}
	colors = append(colors, "blue")
	colors = append(colors, "yellow", "purple")

	moreColors := []string{"black", "white"}
	colors = append(colors, moreColors...) // spread with ...
	fmt.Println("\n5. Final colors slice:", colors)

	// ============================================================
	// 6. SLICES ARE REFERENCE TYPES
	// ============================================================

	original := []int{1, 2, 3}
	reference := original // shares underlying array
	reference[0] = 100
	fmt.Println("\n6. Original changed via reference:", original)

	// ============================================================
	// 7. COPYING A SLICE PROPERLY
	// ============================================================

	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	dst[0] = 999
	fmt.Println("\n7. Source (unchanged):", src)
	fmt.Println("   Independent copy:", dst)

	// ============================================================
	// 8. ITERATING WITH range
	// ============================================================

	fmt.Println("\n8. Iterating:")
	for index, value := range fruits {
		fmt.Printf("   Index %d -> %s\n", index, value)
	}

	// ============================================================
	// 9. SLICE OF SLICES (2D)
	// ============================================================

	grid := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("\n9. Grid:")
	for _, row := range grid {
		fmt.Println("  ", row)
	}

	// ============================================================
	// 10. NIL SLICE vs EMPTY SLICE
	// ============================================================

	var nilSlice []int
	emptySlice := []int{}
	fmt.Println("\n10. nilSlice == nil:", nilSlice == nil)
	fmt.Println("    emptySlice == nil:", emptySlice == nil)

	nilSlice = append(nilSlice, 42) // append works even on nil
	fmt.Println("    After append:", nilSlice)
}

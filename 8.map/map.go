package main

import "fmt"

func main() {
	var m map[string]int     // [key] -> value
	m = make(map[string]int) // make is used to create a map other wise it will be nil
	m["key1"] = 1
	m["key2"] = 2
	fmt.Println(m)

	// the other way to declare and initialize a map is
	map1 := make(map[string]string)
	map1["key1"] = "value1"
	map1["key2"] = "value2"
	fmt.Println(map1)

	// delete is used to remove a key-value pair from a map
	delete(map1, "key1")
	fmt.Println(map1)

	// clear is used to remove all key-value pairs from a map
	clear(map1)
	fmt.Println(map1) // this will print an empty map
}

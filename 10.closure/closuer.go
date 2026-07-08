package main

import "fmt"

func counter() func() int {
    count := 0  // This variable is captured by the closure
    return func() int {
        count++   // Inner function "remembers" count
        return count
    }
}

func main() {
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
    fmt.Println(c()) // 3
    
    d := counter()   // Creates a new counter with its own count
    fmt.Println(d()) // 1 (starts fresh)
}